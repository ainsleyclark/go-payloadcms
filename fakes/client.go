package payloadfakes

import (
	"context"
	"net/http"

	"github.com/ainsleydev/go-payloadcms"
)

// MockService is a mock implementation of the Service interface.
type MockService struct {
	DoFunc            func(ctx context.Context, method, path string, body any, v any) (payloadcms.Response, error)
	DoWithRequestFunc func(ctx context.Context, req *http.Request, v any) (payloadcms.Response, error)
	GetFunc           func(ctx context.Context, path string, v any) (payloadcms.Response, error)
	PostFunc          func(ctx context.Context, path string, in any) (payloadcms.Response, error)
	PutFunc           func(ctx context.Context, path string, in any) (payloadcms.Response, error)
	DeleteFunc        func(ctx context.Context, path string, v any) (payloadcms.Response, error)
}

// NewMockService creates a new fake service stub.
func NewMockService() *MockService {
	return &MockService{
		DoFunc: func(_ context.Context, _ string, _ string, _ any, _ any) (payloadcms.Response, error) {
			return payloadcms.Response{}, nil
		},
		DoWithRequestFunc: func(_ context.Context, _ *http.Request, _ any) (payloadcms.Response, error) {
			return payloadcms.Response{}, nil
		},
		GetFunc: func(_ context.Context, _ string, _ any) (payloadcms.Response, error) {
			return payloadcms.Response{}, nil
		},
		PostFunc: func(_ context.Context, _ string, _ any) (payloadcms.Response, error) {
			return payloadcms.Response{}, nil
		},
		PutFunc: func(_ context.Context, _ string, _ any) (payloadcms.Response, error) {
			return payloadcms.Response{}, nil
		},
		DeleteFunc: func(_ context.Context, _ string, _ any) (payloadcms.Response, error) {
			return payloadcms.Response{}, nil
		},
	}
}

// Do calls the mock implementation.
func (m *MockService) Do(ctx context.Context, method, path string, body any, v any) (payloadcms.Response, error) {
	return m.DoFunc(ctx, method, path, body, v)
}

// DoWithRequest calls the mock implementation.
func (m *MockService) DoWithRequest(ctx context.Context, req *http.Request, v any) (payloadcms.Response, error) {
	return m.DoWithRequestFunc(ctx, req, v)
}

// Get calls the mock implementation.
func (m *MockService) Get(ctx context.Context, path string, v any) (payloadcms.Response, error) {
	return m.GetFunc(ctx, path, v)
}

// Post calls the mock implementation.
func (m *MockService) Post(ctx context.Context, path string, in any) (payloadcms.Response, error) {
	return m.PostFunc(ctx, path, in)
}

// Put calls the mock implementation.
func (m *MockService) Put(ctx context.Context, path string, in any) (payloadcms.Response, error) {
	return m.PutFunc(ctx, path, in)
}

// Delete calls the mock implementation.
func (m *MockService) Delete(ctx context.Context, path string, v any) (payloadcms.Response, error) {
	return m.DeleteFunc(ctx, path, v)
}
