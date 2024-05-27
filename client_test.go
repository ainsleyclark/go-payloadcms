package payloadcms

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	defaultBody     = []byte(`{"id": 1, "name": "John Doe"}`)
	defaultResponse = Entity{ID: 1, Name: "John Doe"}
	defaultHandler  = func(t *testing.T) http.HandlerFunc {
		t.Helper()

		return func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, err := w.Write(defaultBody)
			if err != nil {
				t.Errorf("failed to write response: %v", err)
			}
		}
	}
)

type Entity struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Global test helper functions

func AssertEqual(t *testing.T, expected, actual any) {
	if expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func AssertNoError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func Setup(t *testing.T, handlerFunc http.HandlerFunc) (*Client, func()) {
	t.Helper()

	server := httptest.NewServer(handlerFunc)
	return &Client{
			baseURL: server.URL,
			client:  server.Client(),
			reader:  io.ReadAll,
		}, func() {
			server.Close()
		}
}

func TestNew(t *testing.T) {
	t.Parallel()

	t.Run("OK", func(t *testing.T) {
		t.Parallel()
		_, err := New(WithBaseURL("http://localhost:8080"))
		if err != nil {
			t.Errorf("expected no error: %v", err)
		}
	})

	t.Run("Failed validation", func(t *testing.T) {
		t.Parallel()
		_, err := New()
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("Assigns client", func(t *testing.T) {
		t.Parallel()
		c, err := New(
			WithClient(nil),
			WithBaseURL("http://localhost:8080"),
		)
		if err != nil {
			t.Errorf("expected no error: %v", err)
		}
		if c.client == nil {
			t.Errorf("expected client to be assigned")
		}
	})
}

func TestClientDo(t *testing.T) {
	tt := map[string]struct {
		method   string
		path     string
		wantCode int
		wantBody []byte
		wantErr  bool
	}{
		"Bad request": {
			method:  "INVALID",
			path:    "@£$%&*()",
			wantErr: true,
		},
		"Do error": {
			path:    "@£$%&*()",
			wantErr: true,
		},
		"200 OK": {
			method:   http.MethodGet,
			path:     "/users/1",
			wantCode: http.StatusOK,
			wantBody: defaultBody,
			wantErr:  false,
		},
		//"404 Not Found": {
		//	method:   http.MethodGet,
		//	path:     "/nonexistent",
		//	wantCode: http.StatusNotFound,
		//	wantBody: defaultBody,
		//	wantErr:  true,
		//},
		//"500 Internal Server Error": {
		//	method:   http.MethodGet,
		//	path:     "/error",
		//	wantCode: http.StatusInternalServerError,
		//	wantBody: defaultBody,
		//	wantErr:  true,
		//},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			client, teardown := Setup(t, defaultHandler(t))
			defer teardown()

			response, err := client.Do(context.TODO(), test.method, test.path, nil, nil)

			if test.wantErr != (err != nil) {
				t.Errorf("expected error: %v, got: %v", test.wantErr, err != nil)
			}
			if response.StatusCode != test.wantCode {
				t.Errorf("expected status code: %d, got: %d", test.wantCode, response.StatusCode)
			}
			if string(response.Content) != string(test.wantBody) {
				t.Errorf("expected body: %s, got: %s", string(test.wantBody), string(response.Content))
			}
		})
	}
}

func TestClient_NewRequest(t *testing.T) {
	t.Skip("TODO: implement")
}

func TestErrors_Error(t *testing.T) {
	err := Errors{
		{Message: "error 1"},
		{Message: "error 2"},
	}
	if err.Error() != "error 1, error 2" {
		t.Errorf("expected error message: %s, got: %s", "error 1, error 2", err.Error())
	}
}
