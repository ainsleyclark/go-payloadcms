package payloadfakes

import (
	"context"

	"github.com/ainsleyclark/go-payloadcms"
)

// MockCollectionService is a mock implementation of the CollectionService interface.
type MockCollectionService struct {
	FindByIDFunc   func(ctx context.Context, collection payloadcms.Collection, id int, out any, opts ...payloadcms.RequestOption) (payloadcms.Response, error)
	FindBySlugFunc func(ctx context.Context, collection payloadcms.Collection, slug string, out any, opts ...payloadcms.RequestOption) (payloadcms.Response, error)
	ListFunc       func(ctx context.Context, collection payloadcms.Collection, params payloadcms.ListParams, out any, opts ...payloadcms.RequestOption) (payloadcms.Response, error)
	CreateFunc     func(ctx context.Context, collection payloadcms.Collection, in any, opts ...payloadcms.RequestOption) (payloadcms.Response, error)
	UpdateByIDFunc func(ctx context.Context, collection payloadcms.Collection, id int, in any, opts ...payloadcms.RequestOption) (payloadcms.Response, error)
	DeleteByIDFunc func(ctx context.Context, collection payloadcms.Collection, id int, opts ...payloadcms.RequestOption) (payloadcms.Response, error)
}

// NewMockCollectionService creates a new fake collections stub.
func NewMockCollectionService() *MockCollectionService {
	return &MockCollectionService{
		FindByIDFunc: func(_ context.Context, _ payloadcms.Collection, _ int, _ any, _ ...payloadcms.RequestOption) (payloadcms.Response, error) {
			return payloadcms.Response{}, nil
		},
		FindBySlugFunc: func(_ context.Context, _ payloadcms.Collection, _ string, _ any, _ ...payloadcms.RequestOption) (payloadcms.Response, error) {
			return payloadcms.Response{}, nil
		},
		ListFunc: func(_ context.Context, _ payloadcms.Collection, _ payloadcms.ListParams, _ any, _ ...payloadcms.RequestOption) (payloadcms.Response, error) {
			return payloadcms.Response{}, nil
		},
		CreateFunc: func(_ context.Context, _ payloadcms.Collection, _ any, _ ...payloadcms.RequestOption) (payloadcms.Response, error) {
			return payloadcms.Response{}, nil
		},
		UpdateByIDFunc: func(_ context.Context, _ payloadcms.Collection, _ int, _ any, _ ...payloadcms.RequestOption) (payloadcms.Response, error) {
			return payloadcms.Response{}, nil
		},
		DeleteByIDFunc: func(_ context.Context, _ payloadcms.Collection, _ int, _ ...payloadcms.RequestOption) (payloadcms.Response, error) {
			return payloadcms.Response{}, nil
		},
	}
}

// FindByID calls the mock implementation.
func (m *MockCollectionService) FindByID(ctx context.Context, collection payloadcms.Collection, id int, out any, opts ...payloadcms.RequestOption) (payloadcms.Response, error) {
	return m.FindByIDFunc(ctx, collection, id, out, opts...)
}

// FindBySlug calls the mock implementation.
func (m *MockCollectionService) FindBySlug(ctx context.Context, collection payloadcms.Collection, slug string, out any, opts ...payloadcms.RequestOption) (payloadcms.Response, error) {
	return m.FindBySlugFunc(ctx, collection, slug, out, opts...)
}

// List calls the mock implementation.
func (m *MockCollectionService) List(ctx context.Context, collection payloadcms.Collection, params payloadcms.ListParams, out any, opts ...payloadcms.RequestOption) (payloadcms.Response, error) {
	return m.ListFunc(ctx, collection, params, out, opts...)
}

// Create calls the mock implementation.
func (m *MockCollectionService) Create(ctx context.Context, collection payloadcms.Collection, in any, opts ...payloadcms.RequestOption) (payloadcms.Response, error) {
	return m.CreateFunc(ctx, collection, in, opts...)
}

// UpdateByID calls the mock implementation.
func (m *MockCollectionService) UpdateByID(ctx context.Context, collection payloadcms.Collection, id int, in any, opts ...payloadcms.RequestOption) (payloadcms.Response, error) {
	return m.UpdateByIDFunc(ctx, collection, id, in, opts...)
}

// DeleteByID calls the mock implementation.
func (m *MockCollectionService) DeleteByID(ctx context.Context, collection payloadcms.Collection, id int, opts ...payloadcms.RequestOption) (payloadcms.Response, error) {
	return m.DeleteByIDFunc(ctx, collection, id, opts...)
}
