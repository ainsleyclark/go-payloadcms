package payloadcms

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/gabriel-vasile/mimetype"
)

// MediaService is an interface for uploading media to the Payload API.
// Experimental feature.
//
// See: https://payloadcms.com/docs/upload/overview
type MediaService interface {
	Upload(ctx context.Context, r io.Reader, in, out any, opts MediaOptions) (Response, error)
	UploadFromURL(ctx context.Context, url string, in, out any, opts MediaOptions) (Response, error)
}

// MediaServiceOp represents a service for managing media within Payload.
type MediaServiceOp struct {
	Client *Client
}

// MediaOptions represents non-required options for uploading media.
type MediaOptions struct {
	// The collection to upload the media to, defaults to "media"
	Collection string
	// If set, the file name of the media will be overridden from the file name.
	// Note, this will not change the file extension.
	FileNameOverride string
}

// Upload uploads a file to the media endpoint.
func (s MediaServiceOp) Upload(ctx context.Context, r io.Reader, in, out any, opts MediaOptions) (Response, error) {
	values, err := fileUploadValues(r, in)
	if err != nil {
		return Response{}, err
	}
	return s.uploadFile(ctx, values, out, opts)
}

func (s MediaServiceOp) UploadFromURL(ctx context.Context, url string, in, out any, opts MediaOptions) (Response, error) {
	// Download the file from the URL
	resp, err := s.Client.client.Get(url)
	if err != nil {
		return Response{}, err
	}
	defer resp.Body.Close()

	// Check if the response status is OK
	if resp.StatusCode != http.StatusOK {
		return Response{}, fmt.Errorf("failed to download file: status code %d", resp.StatusCode)
	}

	// Create a temporary file to store the downloaded content
	tmpfile, err := os.Create(filepath.Join(os.TempDir(), fileNameFromURL(url)))
	if err != nil {
		return Response{}, fmt.Errorf("failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpfile.Name()) // Clean up the temporary file

	// Write the downloaded content to the temporary file
	_, err = io.Copy(tmpfile, resp.Body)
	if err != nil {
		return Response{}, fmt.Errorf("failed to write to temporary file: %v", err)
	}

	values, err := fileUploadValues(tmpfile, in)
	if err != nil {
		return Response{}, err
	}

	return s.uploadFile(ctx, values, out, opts)
}

// uploadFile prepares a multipart form and performs the upload request
//   - Takes the context, collection name, map of form values (including the file), and optional output struct
//   - Returns a Response object and any errors encountered
func (s MediaServiceOp) uploadFile(ctx context.Context, values map[string]io.Reader, out any, opts MediaOptions) (Response, error) {
	if opts.Collection == "" {
		opts.Collection = "media"
	}

	// Prepare a multipart form
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	for key, r := range values {
		if x, ok := r.(io.Closer); ok {
			defer x.Close()
		}
		if x, ok := r.(*os.File); ok {
			if err := handleFileUpload(w, key, x, opts); err != nil {
				return Response{}, err
			}
		} else {
			// Add other fields
			fw, err := w.CreateFormField(key)
			if err != nil {
				return Response{}, err
			}
			if _, err := io.Copy(fw, r); err != nil {
				return Response{}, err
			}
		}
	}

	// Close the multipart writer
	if err := w.Close(); err != nil {
		return Response{}, fmt.Errorf("failed to close multipart writer: %v", err)
	}

	p := fmt.Sprintf("/api/%s", opts.Collection)
	req, err := s.Client.NewFormRequest(ctx, http.MethodPost, p, &b, w.FormDataContentType())
	if err != nil {
		return Response{}, err
	}

	return s.Client.DoWithRequest(ctx, req, out)
}

func fileUploadValues(r io.Reader, v any) (map[string]io.Reader, error) {
	if r == nil {
		return nil, fmt.Errorf("file is required")
	}

	// Marshal the `in` struct to JSON for the _payload field
	payloadJSON, err := json.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal input struct: %v", err)
	}

	values := map[string]io.Reader{
		"file":     r,
		"_payload": strings.NewReader(string(payloadJSON)), // The Payload CMS structure
	}

	return values, nil
}

// handleFileUpload adds a file to the multipart writer.
func handleFileUpload(w *multipart.Writer, key string, f *os.File, opts MediaOptions) error {
	// Read the first 512 bytes to detect the MIME type.
	mime, err := mimetype.DetectReader(f) // Only tested with mimetype.DetectFile
	if err != nil {
		return err
	}

	// Reset the file pointer back to the beginning after MIME detection
	if _, err := f.Seek(0, io.SeekStart); err != nil {
		return fmt.Errorf("failed to reset file pointer: %v", err)
	}

	fileName := f.Name()
	if opts.FileNameOverride != "" {
		fileName = opts.FileNameOverride + "." + filepath.Ext(f.Name())[1:]
	}

	// Create a new form part
	fw, err := w.CreatePart(textproto.MIMEHeader{
		"Content-Type": {
			mime.String(),
		},
		"Content-Disposition": {
			fmt.Sprintf(`form-data; name="%s"; filename="%s"`, key, fileName),
		},
	})
	if err != nil {
		return err
	}

	// Copy the remaining file contents to the form part
	_, err = io.Copy(fw, f)
	return err
}

func fileNameFromURL(url string) string {
	parts := strings.Split(url, "/")

	// Get the last part of the URL which contains the filename
	filenameWithExtension := parts[len(parts)-1]

	// Check if the last part contains a dot indicating an extension
	if strings.Contains(filenameWithExtension, ".") {
		// Extract the filename from filenameWithExtension
		filename := path.Base(filenameWithExtension)
		return filename
	}

	// If no dot is found, return an empty string
	return ""
}
