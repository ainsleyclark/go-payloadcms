package payloadfakes

import (
	"context"
)

// MockGlobalsService is a mock implementation of the GlobalsService interface.
type MockGlobalsService struct {
	GetFunc    func(ctx context.Context, global payloadcms.Global, in any) (payloadcms.Response, error)
	UpdateFunc func(ctx context.Context, global payloadcms.Global, in any) (payloadcms.Response, error)
}

// NewMockGlobalsService creates a new fake globals stub.
func NewMockGlobalsService() *MockGlobalsService {
	return &MockGlobalsService{
		GetFunc: func(_ context.Context, _ payloadcms.Global, _ any) (payloadcms.Response, error) {
			return payloadcms.Response{}, nil
		},
		UpdateFunc: func(_ context.Context, _ payloadcms.Global, _ any) (payloadcms.Response, error) {
			return payloadcms.Response{}, nil
		},
	}
}

// Get calls the mock implementation.
func (m *MockGlobalsService) Get(ctx context.Context, global payloadcms.Global, in any) (payloadcms.Response, error) {
	return m.GetFunc(ctx, global, in)
}

// Update calls the mock implementation.
func (m *MockGlobalsService) Update(ctx context.Context, global payloadcms.Global, in any) (payloadcms.Response, error) {
	return m.UpdateFunc(ctx, global, in)
}
