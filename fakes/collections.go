package payloadfakes

import (
	"context"
	"errors"

	"github.com/ainsleydev/webkit/pkg/apis/payloadcms"
)

// MockCollectionService is a mock implementation of the CollectionService interface.
type MockCollectionService struct {
	FindByIdFunc   func(ctx context.Context, collection payloadcms.Collection, id int, out any) (Response, error)
	FindBySlugFunc func(ctx context.Context, collection payloadcms.Collection, slug string, out any) (Response, error)
	ListFunc       func(ctx context.Context, collection payloadcms.Collection, params payloadcms.ListParams, out any) (Response, error)
	CreateFunc     func(ctx context.Context, collection payloadcms.Collection, in any) (Response, error)
	UpdateByIDFunc func(ctx context.Context, collection payloadcms.Collection, id int, in any) (Response, error)
	DeleteByIDFunc func(ctx context.Context, collection payloadcms.Collection, id int) (Response, error)
}

// NewMockCollectionService creates a new MockCollectionService with dpayloadcms.
func NewMockCollectionService() *MockCollectionService {
	return &MockCollectionService{
		FindByIdFunc: func(ctx context.Context, collection payloadcms.Collection, id int, out any) (Response, error) {
			return Response{}, nil
		},
		FindBySlugFunc: func(ctx context.Context, collection payloadcms.Collection, slug string, out any) (Response, error) {
			return Response{}, nil
		},
		ListFunc: func(ctx context.Context, collection payloadcms.Collection, params payloadcms.ListParams, out any) (Response, error) {
			return Response{}, nil
		},
		CreateFunc: func(ctx context.Context, collection payloadcms.Collection, in any) (Response, error) {
			return Response{}, nil
		},
		UpdateByIDFunc: func(ctx context.Context, collection payloadcms.Collection, id int, in any) (Response, error) {
			return Response{}, nil
		},
		DeleteByIDFunc: func(ctx context.Context, collection payloadcms.Collection, id int) (Response, error) {
			return Response{}, nil
		},
	}
}

// FindById calls the mock implementation.
func (m *MockCollectionService) FindById(ctx context.Context, collection payloadcms.Collection, id int, out any) (Response, error) {
	return m.FindByIdFunc(ctx, collection, id, out)
}

// FindBySlug calls the mock implementation.
func (m *MockCollectionService) FindBySlug(ctx context.Context, collection payloadcms.Collection, slug string, out any) (Response, error) {
	return m.FindBySlugFunc(ctx, collection, slug, out)
}

// List calls the mock implementation.
func (m *MockCollectionService) List(ctx context.Context, collection payloadcms.Collection, params payloadcms.ListParams, out any) (Response, error) {
	return m.ListFunc(ctx, collection, params, out)
}

// Create calls the mock implementation.
func (m *MockCollectionService) Create(ctx context.Context, collection payloadcms.Collection, in any) (Response, error) {
	return m.CreateFunc(ctx, collection, in)
}

// UpdateByID calls the mock implementation.
func (m *MockCollectionService) UpdateByID(ctx context.Context, collection payloadcms.Collection, id int, in any) (Response, error) {
	return m.UpdateByIDFunc(ctx, collection, id, in)
}

// DeleteByID calls the mock implementation.
func (m *MockCollectionService) DeleteByID(ctx context.Context, collection payloadcms.Collection, id int) (Response, error) {
	return m.DeleteByIDFunc(ctx, collection, id)
}

// Response is a mock structure to hold the response data.
type Response struct {
	Content []byte
}

// Example usage of the MockCollectionService:
func ExampleUsage() {
	mockService := NewMockCollectionService()

	mockService.FindByIdFunc = func(ctx context.Context, collection payloadcms.Collection, id int, out any) (Response, error) {
		// Custom behavior for FindById
		if id == 1 {
			return Response{Content: []byte(`{"id":1,"name":"Test User"}`)}, nil
		}
		return Response{}, errors.New("not found")
	}

	// Now you can use mockService in your tests
}
