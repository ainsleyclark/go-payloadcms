package payloadcms

import (
	"context"
	"net/http"
	"testing"
)

func TestGlobalsService_Get(t *testing.T) {
	client, teardown := Setup(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write(defaultBody)
		AssertNoError(t, err)
		AssertEqual(t, "/api/globals/settings", r.URL.Path)
		AssertEqual(t, http.MethodGet, r.Method)
	})
	defer teardown()

	s := GlobalsServiceOp{
		Client: client,
	}

	var e Entity
	_, err := s.Get(context.TODO(), "settings", &e)
	AssertNoError(t, err)
	AssertEqual(t, e, defaultResponse)
}

func TestGlobalsService_Update(t *testing.T) {

}
