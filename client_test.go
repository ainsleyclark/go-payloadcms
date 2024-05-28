package payloadcms

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-querystring/query"
)

var (
	defaultBody     = []byte(`{"id": 1, "name": "John Doe"}`)
	defaultResource = Resource{ID: 1, Name: "John Doe"}
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

type Resource struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func Setup(t *testing.T, handlerFunc http.HandlerFunc) (*Client, func()) {
	t.Helper()

	server := httptest.NewServer(handlerFunc)
	return &Client{
			baseURL:     server.URL,
			client:      server.Client(),
			reader:      io.ReadAll,
			queryValues: query.Values,
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
		body     any
		wantCode int
		wantBody []byte
		wantErr  bool
	}{
		"Marshal error": {
			body:    make(chan int),
			wantErr: true,
		},
		"Bad request": {
			method:  "INVALID",
			path:    "@£$%&*()",
			wantErr: true,
		},
		"Do error": {
			path:    "wrong",
			method:  "H",
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

			response, err := client.Do(context.TODO(), test.method, test.path, test.body, nil)
			AssertEqual(t, test.wantErr, err != nil)
			AssertEqual(t, test.wantCode, response.StatusCode)
			AssertEqual(t, string(test.wantBody), string(response.Content))
		})
	}
}

func TestClient_Requests(t *testing.T) {
	path := "/users/1"

	tt := map[string]struct {
		call       func(s *Client) (Response, error)
		wantMethod string
	}{
		"Get": {
			call: func(s *Client) (Response, error) {
				return s.Get(context.TODO(), path, nil)
			},
			wantMethod: http.MethodGet,
		},
		"Post": {
			call: func(s *Client) (Response, error) {
				return s.Post(context.TODO(), path, nil)
			},
			wantMethod: http.MethodPost,
		},
		"Put": {
			call: func(s *Client) (Response, error) {
				return s.Put(context.TODO(), path, nil)
			},
			wantMethod: http.MethodPut,
		},
		"Delete": {
			call: func(s *Client) (Response, error) {
				return s.Delete(context.TODO(), path, nil)
			},
			wantMethod: http.MethodDelete,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			client, teardown := Setup(t, func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				AssertEqual(t, path, r.URL.Path)
				AssertEqual(t, test.wantMethod, r.Method)
			})
			defer teardown()

			_, err := test.call(client)
			AssertNoError(t, err)
		})
	}
}

func TestClient_NewRequest(t *testing.T) {
	c := Client{apiKey: "123"}

	t.Run("OK", func(t *testing.T) {
		got, err := c.NewRequest(context.TODO(), http.MethodGet, "/users/1", nil)
		AssertNoError(t, err)
		AssertEqual(t, http.MethodGet, got.Method)
		AssertEqual(t, "application/json", got.Header.Get("Content-Type"))
		AssertEqual(t, "users API-Key 123", got.Header.Get("Authorization"))
	})

	t.Run("Error", func(t *testing.T) {
		_, err := c.NewRequest(context.TODO(), http.MethodGet, "@£$%", nil)
		AssertError(t, err)
	})
}

func TestClient_NewFormRequest(t *testing.T) {
	c := Client{apiKey: "123"}
	got, err := c.NewFormRequest(context.TODO(), http.MethodGet, "/users/1", nil, "multipart/form-data")
	AssertNoError(t, err)
	AssertEqual(t, http.MethodGet, got.Method)
	AssertEqual(t, "multipart/form-data", got.Header.Get("Content-Type"))
	AssertEqual(t, "users API-Key 123", got.Header.Get("Authorization"))
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
