package payloadcms

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
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
		"FindByStrID": {
			call: func(s CollectionService) (Response, error) {
				return s.FindByStrID(context.Background(), collection, "1", nil)
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
					Where: Query().Equals("colour", "yellow"),
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
			wantMethod: http.MethodPatch,
		},
		"UpdateByStrID": {
			call: func(s CollectionService) (Response, error) {
				return s.UpdateByStrID(context.Background(), collection, "1", defaultResource)
			},
			wantURL:    "/api/posts/1",
			wantMethod: http.MethodPatch,
		},
		"DeleteByID": {
			call: func(s CollectionService) (Response, error) {
				return s.DeleteByID(context.Background(), collection, 1)
			},
			wantURL:    "/api/posts/1",
			wantMethod: http.MethodDelete,
		},
		"DeleteByStrID": {
			call: func(s CollectionService) (Response, error) {
				return s.DeleteByStrID(context.Background(), collection, "1")
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
}

func TestListParams_Encode(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		input ListParams
		want  string
	}{
		"All fields set": {
			input: ListParams{
				Sort:  "name",
				Where: Query().Equals("colour", "yellow"),
				Limit: 10,
				Page:  2,
			},
			want: "?sort=name&field=value&where%5Bcolour%5D%5Bequals%5D=yellow&limit=10&page=2",
		},
		"Only Sort set": {
			input: ListParams{
				Sort: "name",
			},
			want: "?sort=name",
		},
		"Only Limit and Page set": {
			input: ListParams{
				Limit: 5,
				Page:  3,
			},
			want: "?limit=5&page=3",
		},
		"No fields set": {
			input: ListParams{},
			want:  "",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if test.input.Where != nil {
				test.input.Where.params.Add("field", "value")
			}
			got := test.input.Encode()
			assert.Equal(t, test.want, got)
		})
	}
}
