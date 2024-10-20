package payloadcms

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
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
	// Required filename for the upload, you do not need to pass the
	// extension here.
	// Note, this will not change the file extension.
	FileName string
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

	// If filename not provided in options, try to get it from URL
	if opts.FileName == "" {
		filename := fileNameFromURL(url)
		if filename == "" {
			return Response{}, errors.New("no filename provided and couldn't extract from URL")
		}
		// Strip the extension as handleReaderUpload will add the correct one
		opts.FileName = strings.TrimSuffix(filename, filepath.Ext(filename))
	}

	values, err := fileUploadValues(resp.Body, in)
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
		if key == "file" {
			if err := handleReaderUpload(w, key, r, opts.FileName); err != nil {
				return Response{}, err
			}
		} else {
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

// / handleReaderUpload adds a reader's content to the multipart writer with proper MIME type detection
func handleReaderUpload(w *multipart.Writer, key string, r io.Reader, fileName string) error {
	// Buffer the beginning of the file for MIME detection
	var buf bytes.Buffer
	tee := io.TeeReader(r, &buf)

	// Detect MIME type
	mime, err := mimetype.DetectReader(tee)
	if err != nil {
		return fmt.Errorf("failed to detect mime type: %v", err)
	}

	// If no filename is provided, generate one with the correct extension
	if fileName == "" {
		return errors.New("no filename provided")
	}

	// Check if the filename already has an extension
	if ext := filepath.Ext(fileName); ext != "" {
		return fmt.Errorf("filename should not include extension, got: %s", ext)
	}

	fileName = fileName + mime.Extension()

	// Create the form part with the detected MIME type
	h := textproto.MIMEHeader{}
	h.Set("Content-Type", mime.String())
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, key, fileName))

	fw, err := w.CreatePart(h)
	if err != nil {
		return err
	}

	// Write the buffered content first
	if _, err := io.Copy(fw, &buf); err != nil {
		return err
	}

	// Then write the rest of the content
	_, err = io.Copy(fw, r)
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
