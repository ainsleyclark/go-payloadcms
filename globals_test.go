package payloadcms

import (
	"context"
	"net/http"
	"testing"
)

func TestGlobalsService(t *testing.T) {
	global := Global("settings")

	tt := map[string]struct {
		call       func(s GlobalsService) (Response, error)
		wantMethod string
	}{
		"Get": {
			call: func(s GlobalsService) (Response, error) {
				return s.Get(context.Background(), global, nil)
			},
			wantMethod: http.MethodGet,
		},
		"Update": {
			call: func(s GlobalsService) (Response, error) {
				return s.Update(context.Background(), global, nil)
			},
			wantMethod: http.MethodPost,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			client, teardown := Setup(t, func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				_, err := w.Write(defaultBody)
				AssertNoError(t, err)
				AssertEqual(t, "/api/globals/"+string(global), r.URL.Path)
				AssertEqual(t, test.wantMethod, r.Method)
			})
			defer teardown()

			resp, err := test.call(&GlobalsServiceOp{Client: client})
			AssertNoError(t, err)
			AssertEqual(t, string(resp.Content), string(defaultBody))
		})
	}
}
