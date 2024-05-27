package payloadcms

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// GlobalsService is an interface for interacting with the global
// endpoints of the Payload API.
//
// See: https://payloadcms.com/docs/rest-api/overview#globals
type GlobalsService interface {
	Get(ctx context.Context, global Global, in any) (Response, error)
	Update(ctx context.Context, global Global, in any) (Response, error)
}

// GlobalsServiceOp handles communication with the global related
// methods of the Payload API.
type GlobalsServiceOp struct {
	Client *Client
}

// Global represents a global slug from Payload.
type Global string

// Get finds a global by its slug.
func (s GlobalsServiceOp) Get(ctx context.Context, global Global, in any) (Response, error) {
	path := fmt.Sprintf("/api/globals/%s", global)
	return s.Client.Get(ctx, path, in)
}

// Update updates a global by its slug.
func (s GlobalsServiceOp) Update(ctx context.Context, global Global, in any) (Response, error) {
	path := fmt.Sprintf("/api/globals/%s", global)
	buf, err := json.Marshal(in)
	if err != nil {
		return Response{}, err
	}
	return s.Client.Do(ctx, http.MethodPost, path, bytes.NewReader(buf), nil)
}
