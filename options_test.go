package payloadcms

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestClientOptions(t *testing.T) {
	t.Parallel()

	var (
		client  = &http.Client{}
		baseURL = "https://api.payloadcms.com"
		apiKey  = "api-key"
	)

	got, err := New(
		WithClient(client),
		WithBaseURL(baseURL),
		WithAPIKey(apiKey),
	)

	AssertNoError(t, err)
	AssertEqual(t, got.client, client)
	AssertEqual(t, got.baseURL, baseURL)
	AssertEqual(t, got.apiKey, apiKey)
}

func TestRequestOptions(t *testing.T) {
	client, teardown := Setup(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write(defaultBody)
		AssertNoError(t, err)
		query := r.URL.Query()
		assert.Equal(t, "10", query.Get("depth"))
		assert.Equal(t, "value", query.Get("key"))
	})
	defer teardown()

	t.Run("Collections", func(t *testing.T) {
		col := &CollectionServiceOp{Client: client}

		_, err := col.FindByID(context.TODO(), "posts", 1, nil,
			WithDepth(10),
			WithQueryParam("key", "value"),
		)
		require.NoError(t, err)
	})

	t.Run("Globals", func(t *testing.T) {
		col := &GlobalsServiceOp{Client: client}

		_, err := col.Get(context.TODO(), "settings", nil,
			WithDepth(10),
			WithQueryParam("key", "value"),
		)
		require.NoError(t, err)
	})
}
