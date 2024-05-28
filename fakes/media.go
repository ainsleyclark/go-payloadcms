package payloadfakes

import (
	"context"
	"os"
)

// MockMediaService is a mock implementation of the MediaService interface.
type MockMediaService struct {
	UploadFunc        func(ctx context.Context, f *os.File, in, out any, opts payloadcms.MediaOptions) (payloadcms.Response, error)
	UploadFromURLFunc func(ctx context.Context, url string, in, out any, opts payloadcms.MediaOptions) (payloadcms.Response, error)
}

// NewMockMediaService creates a new fake media service stub.
func NewMockMediaService() *MockMediaService {
	return &MockMediaService{
		UploadFunc: func(_ context.Context, _ *os.File, _ any, _ any, _ payloadcms.MediaOptions) (payloadcms.Response, error) {
			return payloadcms.Response{}, nil
		},
		UploadFromURLFunc: func(_ context.Context, _ string, _ any, _ any, _ payloadcms.MediaOptions) (payloadcms.Response, error) {
			return payloadcms.Response{}, nil
		},
	}
}

// Upload calls the mock implementation.
func (m *MockMediaService) Upload(ctx context.Context, f *os.File, in, out any, opts payloadcms.MediaOptions) (payloadcms.Response, error) {
	return m.UploadFunc(ctx, f, in, out, opts)
}

// UploadFromURL calls the mock implementation.
func (m *MockMediaService) UploadFromURL(ctx context.Context, url string, in, out any, opts payloadcms.MediaOptions) (payloadcms.Response, error) {
	return m.UploadFromURLFunc(ctx, url, in, out, opts)
}
