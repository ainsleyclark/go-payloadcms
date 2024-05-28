package payloadcms

import (
	"context"
	"fmt"
	"net/http"
)

// CollectionService is an interface for interacting with the collection
// endpoints of the Payload API.
//
// See: https://payloadcms.com/docs/rest-api/overview#collections
type CollectionService interface {
	FindByID(ctx context.Context, collection Collection, id int, out any) (Response, error)
	FindBySlug(ctx context.Context, collection Collection, slug string, out any) (Response, error)
	List(ctx context.Context, collection Collection, params ListParams, out any) (Response, error)
	Create(ctx context.Context, collection Collection, in any) (Response, error)
	UpdateByID(ctx context.Context, collection Collection, id int, in any) (Response, error)
	DeleteByID(ctx context.Context, collection Collection, id int) (Response, error)
	// TODO: Need to finalise the Delete endpoint which takes in where query params.
}

// CollectionServiceOp handles communication with the collection related
// methods of the Payload API.
type CollectionServiceOp struct {
	Client *Client
}

// Collection represents a collection slug from Payload.
// It's defined as a string under slug within the Collection Config.
type Collection string

const (
	// CollectionUsers defines the Payload users collection slug.
	CollectionUsers Collection = "users"
)

// AllItems is a constant that can be used to retrieve all items from a collection.
// It's defined as 0 in the Payload API.
const AllItems = 0

type (
	// ListParams represents additional query parameters for the find endpoint.
	ListParams struct {
		Sort  string         `json:"sort" url:"sort"`   // Sort the returned documents by a specific field.
		Where map[string]any `json:"where" url:"where"` // Constrain returned documents with a where query.
		Limit int            `json:"limit" url:"limit"` // Limit the returned documents to a certain number.
		Page  int            `json:"page" url:"page"`   // Get a specific page of documents.
	}
	// ListResponse represents a list of entities that is sent back
	// from the Payload CMS.
	ListResponse[T any] struct {
		Docs          []T  `json:"docs"`
		Total         int  `json:"total"`
		TotalDocs     int  `json:"totalDocs"`
		Limit         int  `json:"limit"`
		TotalPages    int  `json:"totalPages"`
		Page          int  `json:"page"`
		PagingCounter int  `json:"pagingCounter"`
		HasPrevPage   bool `json:"hasPrevPage"`
		HasNextPage   bool `json:"hasNextPage"`
		PrevPage      any  `json:"prevPage"`
		NextPage      any  `json:"nextPage"` //no:lint tagliatelle
	}
	// CreateResponse represents a response from the Payload CMS
	// when a new entity is created.
	CreateResponse[T any] struct {
		Doc     T      `json:"doc"`
		Message string `json:"message"`
		Errors  []any  `json:"errors"`
	}
	// UpdateResponse represents a response from the Payload CMS
	// when an entity is updated.
	UpdateResponse[T any] struct {
		Doc     T      `json:"doc"`
		Message string `json:"message"`
		Errors  []any  `json:"error"`
	}
)

// FindByID finds a collection entity by its ID.
func (s CollectionServiceOp) FindByID(ctx context.Context, collection Collection, id int, out any) (Response, error) {
	path := fmt.Sprintf("/api/%s/%d", collection, id)
	return s.Client.Do(ctx, http.MethodGet, path, nil, out)
}

// FindBySlug finds a collection entity by its slug.
// Note: This is not a standard Payload Rest endpoint, but included for convenience.
// If you want to use this endpoint, you'll need to add an express handler
// to your Payload config.
func (s CollectionServiceOp) FindBySlug(ctx context.Context, collection Collection, slug string, out any) (Response, error) {
	path := fmt.Sprintf("/api/%s/%s", collection, slug)
	return s.Client.Do(ctx, http.MethodGet, path, nil, out)
}

// List lists all collection entities.
func (s CollectionServiceOp) List(ctx context.Context, collection Collection, params ListParams, out any) (Response, error) {
	v, err := s.Client.queryValues(params)
	if err != nil {
		return Response{}, err
	}
	path := fmt.Sprintf("/api/%s?%s", collection, v.Encode())
	return s.Client.Do(ctx, http.MethodGet, path, nil, out)
}

// Create creates a new collection entity.
func (s CollectionServiceOp) Create(ctx context.Context, collection Collection, in any) (Response, error) {
	path := fmt.Sprintf("/api/%s", collection)
	return s.Client.Do(ctx, http.MethodPost, path, in, nil)
}

// UpdateByID updates a collection entity by its ID.
func (s CollectionServiceOp) UpdateByID(ctx context.Context, collection Collection, id int, in any) (Response, error) {
	path := fmt.Sprintf("/api/%s/%d", collection, id)
	return s.Client.Do(ctx, http.MethodPut, path, in, nil)
}

// DeleteByID deletes a collection entity by its ID.
func (s CollectionServiceOp) DeleteByID(ctx context.Context, collection Collection, id int) (Response, error) {
	path := fmt.Sprintf("/api/%s/%d", collection, id)
	return s.Client.Do(ctx, http.MethodDelete, path, nil, nil)
}
