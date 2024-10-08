package payloadcms

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"testing"
)

func TestCollectionsService(t *testing.T) {
	t.Parallel()

	collection := Collection("posts")

	tt := map[string]struct {
		call       func(s CollectionService) (Response, error)
		wantURL    string
		wantMethod string
	}{
		"FindByID": {
			call: func(s CollectionService) (Response, error) {
				return s.FindByID(context.Background(), collection, 1, nil)
			},
			wantURL:    "/api/posts/1",
			wantMethod: http.MethodGet,
		},
		"FindBySlug": {
			call: func(s CollectionService) (Response, error) {
				return s.FindBySlug(context.Background(), collection, "slug", nil)
			},
			wantURL:    "/api/posts/slug/slug",
			wantMethod: http.MethodGet,
		},
		"List": {
			call: func(s CollectionService) (Response, error) {
				return s.List(context.Background(), collection, ListParams{
					Sort:  "asc",
					Limit: 10,
					Page:  1,
				}, nil)
			},
			wantURL:    "/api/posts",
			wantMethod: http.MethodGet,
		},
		"Create": {
			call: func(s CollectionService) (Response, error) {
				return s.Create(context.Background(), collection, defaultResource)
			},
			wantURL:    "/api/posts",
			wantMethod: http.MethodPost,
		},
		"UpdateByID": {
			call: func(s CollectionService) (Response, error) {
				return s.UpdateByID(context.Background(), collection, 1, defaultResource)
			},
			wantURL:    "/api/posts/1",
			wantMethod: http.MethodPut,
		},
		"DeleteByID": {
			call: func(s CollectionService) (Response, error) {
				return s.DeleteByID(context.Background(), collection, 1)
			},
			wantURL:    "/api/posts/1",
			wantMethod: http.MethodDelete,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			client, teardown := Setup(t, func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				_, err := w.Write(defaultBody)
				AssertNoError(t, err)
				AssertEqual(t, test.wantURL, r.URL.Path)
				AssertEqual(t, test.wantMethod, r.Method)
			})
			defer teardown()

			resp, err := test.call(&CollectionServiceOp{Client: client})
			AssertNoError(t, err)
			AssertEqual(t, string(resp.Content), string(defaultBody))
		})
	}

	t.Run("List returns error on QueryValues", func(t *testing.T) {
		t.Parallel()

		client, teardown := Setup(t, defaultHandler(t))
		defer teardown()
		client.queryValues = func(_ any) (url.Values, error) {
			return nil, errors.New("query error")
		}
		client.Collections = CollectionServiceOp{Client: client}
		_, err := client.Collections.List(context.Background(), collection, ListParams{}, nil)
		AssertError(t, err)
	})
}
