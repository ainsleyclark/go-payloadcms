package payloadcms

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"
)

// Service is an interface that defines common methods for interacting
// with the Payload API.
//
// See: https://payloadcms.com/docs/rest-api/overview
type Service interface {
	Do(ctx context.Context, method, path string, body any, v any, opts ...RequestOption) (Response, error)
	DoWithRequest(_ context.Context, req *http.Request, v any, opts ...RequestOption) (Response, error)
	Get(ctx context.Context, path string, v any, opts ...RequestOption) (Response, error)
	Post(ctx context.Context, path string, in any, opts ...RequestOption) (Response, error)
	Put(ctx context.Context, path string, in any, opts ...RequestOption) (Response, error)
	Delete(ctx context.Context, path string, v any, opts ...RequestOption) (Response, error)
}

// Client represents a Payload CMS client.
// For more information, see https://payloadcms.com/docs/api.
type Client struct {
	// Each collection is mounted using its slug value. For example, if a collection's slug is
	// users, all corresponding routes will be mounted on /api/users.
	// For more info, visit: https://payloadcms.com/docs/rest-api/overview#collections
	Collections CollectionService

	// Globals cannot be created or deleted, so there are only two REST endpoints opened:
	// For more info, visit: https://payloadcms.com/docs/rest-api/overview#globals
	Globals GlobalsService

	// Media is a separate service used to upload and manage media files.
	// For more info, visit: https://payloadcms.com/docs/upload/overview
	Media MediaService

	// TODO:
	// - Auth:		 	https://payloadcms.com/docs/rest-api/overview#auth-operations
	// - Preferences: 	https://payloadcms.com/docs/rest-api/overview#preferences

	// Private fields
	client      *http.Client
	baseURL     string
	apiKey      string
	reader      func(io.Reader) ([]byte, error)
	queryValues func(v any) (url.Values, error)
}

var _ Service = (*Client)(nil)

// New creates a new Payload CMS client.
func New(options ...ClientOption) (*Client, error) {
	c := &Client{
		client:      http.DefaultClient,
		reader:      io.ReadAll,
		queryValues: query.Values,
	}

	// Apply all the functional options to configure the client.
	for _, opt := range options {
		opt(c)
	}

	// Ensure the client has a base URL and client is configured
	if err := c.validate(); err != nil {
		return nil, err
	}

	// Initialize the services
	c.Collections = CollectionServiceOp{Client: c}
	c.Globals = GlobalsServiceOp{Client: c}
	c.Media = MediaServiceOp{Client: c}

	return c, nil
}

func (c *Client) validate() error {
	if c.client == nil {
		c.client = http.DefaultClient
	}
	if c.baseURL == "" {
		return errors.New("baseURL is required")
	}

	return nil
}

// Response is a PayloadAPI API response. This wraps the standard http.Response
// returned from Payload and provides convenient access to things like
// body bytes.
type Response struct {
	*http.Response
	Content []byte `json:"-"`
	Message string `json:"-"`
	Errors  Errors `json:"errors"`
}

// Errors defines a list of Payload API errors.
// For Example
// { "errors": [ { "message": "You are not allowed to perform this action." } ] }
type Errors []Error

// Error defines a singular API error.
type Error struct {
	Message string `json:"message"`
}

// Error implements the error interface to return the error message.
func (e Errors) Error() string {
	var errs []string
	for _, err := range e {
		errs = append(errs, err.Message)
	}
	return strings.Join(errs, ", ")
}

// Do sends an API request and returns the API response. The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred.
//
// Errors occur in the eventuality if the http.StatusCode is not 2xx.
func (c *Client) Do(ctx context.Context, method, path string, body any, v any, opts ...RequestOption) (Response, error) {
	defR := Response{
		Response: &http.Response{},
	}

	if body == nil {
		body = make(map[string]any)
	}

	buf, err := json.Marshal(body)
	if err != nil {
		return defR, err
	}

	uri := fmt.Sprintf("%s/%s", c.baseURL, strings.TrimPrefix(path, "/"))
	req, err := http.NewRequestWithContext(ctx, method, uri, bytes.NewReader(buf))
	if err != nil {
		return defR, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "users API-Key "+c.apiKey)

	r, err := c.performRequest(req, opts...)
	if err != nil {
		return r, err
	}

	if v == nil {
		return r, nil
	}

	return r, json.Unmarshal(r.Content, v)
}

// DoWithRequest sends an API request using the provided http.Request and returns the API response.
// The API response is JSON decoded and stored in the value pointed to by v, or returned
// as an error if an API error has occurred.
func (c *Client) DoWithRequest(_ context.Context, req *http.Request, v any, opts ...RequestOption) (Response, error) {
	r, err := c.performRequest(req, opts...)
	if err != nil {
		return r, err
	}
	if v == nil {
		return r, nil
	}
	return r, json.Unmarshal(r.Content, v)
}

// Get sends an HTTP GET request and returns the API response.
// The API response is JSON decoded and stored in the value pointed to by v, or returned
// as an error if an API error has occurred.
func (c *Client) Get(ctx context.Context, path string, v any, opts ...RequestOption) (Response, error) {
	return c.Do(ctx, http.MethodGet, path, nil, v, opts...)
}

// Post sends an HTTP POST request and returns the API response.
// The API response is JSON decoded and stored in the value pointed to by v, or returned
// as an error if an API error has occurred.
func (c *Client) Post(ctx context.Context, path string, in any, opts ...RequestOption) (Response, error) {
	return c.Do(ctx, http.MethodPost, path, in, nil, opts...)
}

// Put sends an HTTP PUT request and returns the API response.
// The API response is JSON decoded and stored in the value pointed to by v, or returned
// as an error if an API error has occurred.
func (c *Client) Put(ctx context.Context, path string, in any, opts ...RequestOption) (Response, error) {
	return c.Do(ctx, http.MethodPut, path, in, nil, opts...)
}

// Patch sends an HTTP PATCH request and returns the API response.
// The API response is JSON decoded and stored in the value pointed to by v, or returned
// as an error if an API error has occurred.
func (c *Client) Patch(ctx context.Context, path string, in any, opts ...RequestOption) (Response, error) {
	return c.Do(ctx, http.MethodPatch, path, in, nil, opts...)
}

// Delete sends an HTTP DELETE request and returns the API response.
// The API response is JSON decoded and stored in the value pointed to by v, or returned
// as an error if an API error has occurred.
func (c *Client) Delete(ctx context.Context, path string, v any, opts ...RequestOption) (Response, error) {
	return c.Do(ctx, http.MethodDelete, path, nil, v, opts...)
}

// NewRequest creates a new Payload API request by using the API
// Key within the client. A method, path and body is attached to
// the request.
func (c *Client) NewRequest(ctx context.Context, method, path string, body io.Reader) (*http.Request, error) {
	uri := fmt.Sprintf("%s/%s", c.baseURL, strings.TrimPrefix(path, "/"))
	req, err := http.NewRequestWithContext(ctx, method, uri, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "users API-Key "+c.apiKey)

	return req, nil
}

func (c *Client) NewFormRequest(ctx context.Context, method, path string, body io.Reader, contentType string) (*http.Request, error) {
	req, err := c.NewRequest(ctx, method, path, body)
	if err != nil {
		return nil, err
	}

	// Set the content type to contain the boundary.
	req.Header.Set("Content-Type", contentType)

	return req, nil
}

func (c *Client) performRequest(req *http.Request, opts ...RequestOption) (Response, error) {
	for _, opt := range opts {
		opt(req)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return Response{Response: &http.Response{}}, err
	}
	defer resp.Body.Close()

	r := Response{Response: resp}

	buf, err := c.reader(resp.Body)
	if err != nil {
		return r, err
	}
	r.Content = buf

	if !is2xx(resp.StatusCode) {
		if string(buf) == "" {
			return r, errors.New("received no body with status code: " + resp.Status)
		}
		if err := json.Unmarshal(buf, &r); err != nil {
			return r, errors.New("failed to unmarshal error response: " + err.Error())
		}
		return r, fmt.Errorf("unexpected status code: %d, errors: %v",
			resp.StatusCode,
			r.Errors,
		)
	}

	return r, nil
}

func is2xx(statusCode int) bool {
	return statusCode >= 200 && statusCode < 300
}
