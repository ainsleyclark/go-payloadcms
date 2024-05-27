package payloadcms

import (
	"net/http"
	"testing"
)

func TestOptions(t *testing.T) {
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
