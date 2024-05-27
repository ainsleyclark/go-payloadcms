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

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if got.client != client {
		t.Errorf("expected client %v, got %v", client, got.client)
	}
	if got.baseURL != baseURL {
		t.Errorf("expected baseURL %s, got %s", baseURL, got.baseURL)
	}
	if got.apiKey != apiKey {
		t.Errorf("expected apiKey %s, got %s", apiKey, got.apiKey)
	}
}
