package payloadcms

import (
	"net/http"
	"testing"
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
	t.Parallel()

}
