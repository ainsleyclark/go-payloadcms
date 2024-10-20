package payloadcms

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/google/go-querystring/query"
)

var (
	defaultBody     = []byte(`{"id": 1, "name": "John Doe"}`)
	defaultResource = Resource{
		ID:   1,
		Name: "John Doe",
	}
	defaultHandler = func(t *testing.T) http.HandlerFunc {
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
	t.Parallel()

	t.Run("Marshal error", func(t *testing.T) {
		t.Parallel()

		client, teardown := Setup(t, defaultHandler(t))
		defer teardown()

		resp, err := client.Do(context.TODO(), http.MethodGet, client.baseURL, make(chan int), nil)
		AssertError(t, err)
		AssertEqual(t, resp.Response != nil, true)
	})

	t.Run("New Request Error", func(t *testing.T) {
		t.Parallel()

		client, teardown := Setup(t, defaultHandler(t))
		defer teardown()

		resp, err := client.Do(context.TODO(), http.MethodGet, "@£$%&*()", nil, nil)
		AssertError(t, err)
		AssertEqual(t, resp.Response != nil, true)
	})

	t.Run("Unmarshalls OK", func(t *testing.T) {
		t.Parallel()

		client, teardown := Setup(t, defaultHandler(t))
		defer teardown()

		var r Resource
		resp, err := client.Do(context.TODO(), http.MethodGet, client.baseURL, nil, &r)
		AssertNoError(t, err)
		AssertEqual(t, defaultResource, r)
		AssertEqual(t, string(defaultBody), string(resp.Content))
	})

	t.Run("Read Error", func(t *testing.T) {
		t.Parallel()

		client, teardown := Setup(t, defaultHandler(t))
		defer teardown()
		client.reader = func(_ io.Reader) ([]byte, error) {
			return nil, io.ErrUnexpectedEOF
		}

		_, err := client.Do(context.TODO(), http.MethodGet, client.baseURL, nil, nil)
		AssertError(t, err)
		AssertEqual(t, io.ErrUnexpectedEOF, err)
	})

	t.Run("Bad Request", func(t *testing.T) {
		t.Parallel()

		client, teardown := Setup(t, func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			_, err := w.Write([]byte(`{"errors":[{"message":"You are not allowed to perform this action."}] }`))
			AssertNoError(t, err)
		})
		defer teardown()

		resp, err := client.Do(context.TODO(), http.MethodGet, client.baseURL, nil, nil)
		AssertError(t, err)
		AssertEqual(t, 400, resp.StatusCode)
		AssertEqual(t, "You are not allowed to perform this action.", resp.Errors[0].Message)
		AssertEqual(t, "unexpected status code: 400, errors: You are not allowed to perform this action.", err.Error())
	})

	t.Run("Unmarshal Errors from Response", func(t *testing.T) {
		t.Parallel()

		client, teardown := Setup(t, func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			_, err := w.Write([]byte(`wrong`))
			AssertNoError(t, err)
		})
		defer teardown()

		resp, err := client.Do(context.TODO(), http.MethodGet, client.baseURL, nil, nil)
		AssertError(t, err)
		AssertEqual(t, 400, resp.StatusCode)
		AssertEqual(t, strings.Contains(err.Error(), "failed to unmarshal error response"), true)
	})

	t.Run("No Body", func(t *testing.T) {
		t.Parallel()

		client, teardown := Setup(t, func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusMethodNotAllowed)
		})
		defer teardown()

		_, err := client.Do(context.TODO(), http.MethodGet, client.baseURL, nil, nil)
		AssertError(t, err)
		AssertContains(t, err.Error(), "received no body with status code")
	})
}

func TestClientDoWithRequest(t *testing.T) {
	t.Parallel()

	t.Run("Marshal error", func(t *testing.T) {
		t.Parallel()

		client, teardown := Setup(t, defaultHandler(t))
		defer teardown()

		_, err := client.DoWithRequest(context.TODO(), &http.Request{}, nil)
		AssertError(t, err)
	})

	t.Run("200 OK", func(t *testing.T) {
		t.Parallel()

		client, teardown := Setup(t, defaultHandler(t))
		defer teardown()

		req, err := http.NewRequest(http.MethodGet, client.baseURL, nil)
		AssertNoError(t, err)

		_, err = client.DoWithRequest(context.TODO(), req, nil)
		AssertNoError(t, err)
	})

	t.Run("Unmarshalls OK", func(t *testing.T) {
		t.Parallel()

		client, teardown := Setup(t, defaultHandler(t))
		defer teardown()

		req, err := http.NewRequest(http.MethodGet, client.baseURL, nil)
		AssertNoError(t, err)
		var r Resource
		_, err = client.DoWithRequest(context.TODO(), req, &r)

		AssertNoError(t, err)
		AssertEqual(t, defaultResource, r)
	})
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
	t.Parallel()

	t.Run("OK", func(t *testing.T) {
		t.Parallel()

		c := Client{apiKey: "123"}
		got, err := c.NewRequest(context.TODO(), http.MethodGet, "/users/1", nil)

		AssertNoError(t, err)
		AssertEqual(t, http.MethodGet, got.Method)
		AssertEqual(t, "application/json", got.Header.Get("Content-Type"))
		AssertEqual(t, "users API-Key 123", got.Header.Get("Authorization"))
	})

	t.Run("Error", func(t *testing.T) {
		t.Parallel()

		c := Client{apiKey: "123"}
		_, err := c.NewRequest(context.TODO(), http.MethodGet, "@£$%", nil)
		AssertError(t, err)
	})
}

func TestClient_NewFormRequest(t *testing.T) {
	t.Parallel()

	t.Run("OK", func(t *testing.T) {
		t.Parallel()

		c := Client{apiKey: "123"}
		got, err := c.NewFormRequest(context.TODO(), http.MethodGet, "/users/1", nil, "multipart/form-data")

		AssertNoError(t, err)
		AssertEqual(t, http.MethodGet, got.Method)
		AssertEqual(t, "multipart/form-data", got.Header.Get("Content-Type"))
		AssertEqual(t, "users API-Key 123", got.Header.Get("Authorization"))
	})

	t.Run("Error", func(t *testing.T) {
		t.Parallel()

		c := Client{apiKey: "123"}
		_, err := c.NewFormRequest(context.TODO(), http.MethodGet, "@£$%", nil, "multipart/form-data")
		AssertError(t, err)
	})
}

func TestErrors_Error(t *testing.T) {
	t.Parallel()
	err := Errors{
		{Message: "error 1"},
		{Message: "error 2"},
	}
	if err.Error() != "error 1, error 2" {
		t.Errorf("expected error message: %s, got: %s", "error 1, error 2", err.Error())
	}
}
