package payloadcms

import (
	"net/http"
	"strconv"
)

// ClientOption is a functional option type that allows us to configure the Client.
type ClientOption func(*Client)

// WithClient is a functional option to set the HTTP client of the Payload API.
func WithClient(client *http.Client) ClientOption {
	return func(c *Client) {
		c.client = client
	}
}

// WithBaseURL is a functional option to set the base URL of the Payload API.
// Example: https://api.payloadcms.com
func WithBaseURL(url string) ClientOption {
	return func(c *Client) {
		c.baseURL = url
	}
}

// WithAPIKey is a functional option to set the API key for the Payload API.
// To get an API key, visit: https://payloadcms.com/docs/rest-api/overview#authentication
//
// Usually, you can obtain one by enabling auth on the users type, and
// visiting the users collection in the Payload dashboard.
func WithAPIKey(apiKey string) ClientOption {
	return func(c *Client) {
		c.apiKey = apiKey
	}
}

// requestOptions defines optional parameters for API requests.
type requestOptions struct {
	params map[string]string
}

// RequestOption is a functional option type used to configure request options.
type RequestOption func(*requestOptions)

// WithDepth sets the depth level for API responses.
// Depth determines how much nested data is included in the response.
//
// See: https://payloadcms.com/docs/queries/depth
func WithDepth(depth int) RequestOption {
	return func(c *requestOptions) {
		WithQueryParam("depth", strconv.Itoa(depth))(c)
	}
}

// WithQueryParam adds a query parameter to the API request.
func WithQueryParam(key, val string) RequestOption {
	return func(c *requestOptions) {
		if c.params == nil {
			c.params = make(map[string]string)
		}
		c.params[key] = val
	}
}
