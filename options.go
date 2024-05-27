package payloadcms

import (
	"net/http"
)

// Option is a functional option type that allows us to configure the Client.
type Option func(*Client)

// WithClient is a functional option to set the HTTP client of the Payload API.
func WithClient(client *http.Client) Option {
	return func(c *Client) {
		c.client = client
	}
}

// WithBaseURL is a functional option to set the base URL of the Payload API.
// Example: https://api.payloadcms.com
func WithBaseURL(url string) Option {
	return func(c *Client) {
		c.baseURL = url
	}
}

// WithAPIKey is a functional option to set the API key for the Payload API.
// To get an API key, visit: https://payloadcms.com/docs/rest-api/overview#authentication
//
// Usually, you can obtain one by enabling auth on the users type, and
// visiting the users collection in the Payload dashboard.
func WithAPIKey(apiKey string) Option {
	return func(c *Client) {
		c.apiKey = apiKey
	}
}
