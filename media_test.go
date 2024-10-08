package payloadcms

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"testing"
)

type mediaFields struct {
	Alt     string `json:"alt"`
	Caption string `json:"caption"`
	NoTag   string
}

var mediaData = mediaFields{
	Alt:     "John Doe",
	Caption: "Hello World",
	NoTag:   "No Tag",
}

func createTestFile(t *testing.T, content []byte) (*os.File, func(), error) {
	t.Helper()

	// Create a temporary file
	file, err := os.CreateTemp(t.TempDir(), "testfile.txt")
	if err != nil {
		return nil, func() {}, err
	}

	teardown := func() {
		AssertNoError(t, os.Remove(file.Name()))
	}

	// Write content to the file
	_, err = file.Write(content)
	if err != nil {
		return nil, teardown, err
	}

	// Close the file to flush the content to disk
	err = file.Close()
	if err != nil {
		return nil, teardown, err
	}

	// Open the file for reading and return the os.File pointer
	f, err := os.Open(file.Name())
	return f, teardown, err
}

func TestMediaService_Upload(t *testing.T) {
	t.Parallel()

	t.Run("Nil File", func(t *testing.T) {
		t.Parallel()

		client, teardown := Setup(t, defaultHandler(t))
		defer teardown()

		m := MediaServiceOp{Client: client}
		_, err := m.Upload(context.TODO(), nil, nil, nil, MediaOptions{})
		AssertError(t, err)
	})

	t.Run("OK", func(t *testing.T) {
		t.Parallel()

		file, clean, err := createTestFile(t, []byte("Payload File"))
		defer clean()
		AssertNoError(t, err)

		client, teardown := Setup(t, defaultHandler(t))
		defer teardown()

		m := MediaServiceOp{Client: client}
		r, err := m.Upload(context.TODO(), file, mediaData, nil, MediaOptions{})
		AssertNoError(t, err)
		AssertEqual(t, string(r.Content), string(defaultBody))
	})

	t.Run("Client Error", func(t *testing.T) {
		t.Parallel()

		file, clean, err := createTestFile(t, []byte("Payload File"))
		defer clean()
		AssertNoError(t, err)

		client, teardown := Setup(t, func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		})
		defer teardown()

		m := MediaServiceOp{Client: client}
		_, err = m.Upload(context.TODO(), file, mediaData, nil, MediaOptions{})
		AssertError(t, err)
	})
}

func TestMediaService_UploadFromURL(t *testing.T) {
	t.Parallel()

	t.Run("Client Error", func(t *testing.T) {
		t.Parallel()

		client, teardown := Setup(t, func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		})
		defer teardown()

		m := MediaServiceOp{Client: client}
		_, err := m.UploadFromURL(context.TODO(), "https://example.com", nil, nil, MediaOptions{})
		AssertError(t, err)
	})

	t.Run("OK", func(t *testing.T) {
		t.Parallel()

		client, teardown := Setup(t, func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusOK)
			_, err := w.Write(defaultBody)
			AssertNoError(t, err)
		})
		defer teardown()

		m := MediaServiceOp{Client: client}
		r, err := m.UploadFromURL(context.TODO(), "https://example.com", nil, nil, MediaOptions{})
		AssertNoError(t, err)
		AssertEqual(t, string(r.Content), string(defaultBody))
	})
}

func TestGetUploadValues(t *testing.T) {
	t.Parallel()

	t.Run("File is required", func(t *testing.T) {
		t.Parallel()
		_, err := fileUploadValues(nil, mediaData)
		AssertEqual(t, true, err != nil)
	})

	t.Run("Struct with JSON tags", func(t *testing.T) {
		t.Parallel()

		f, teardown, err := createTestFile(t, []byte("Payload File"))
		defer teardown()
		AssertNoError(t, err)

		got, err := fileUploadValues(f, mediaData)
		AssertEqual(t, false, err != nil)
		AssertEqual(t, 2, len(got))

		// Test file content
		file, err := io.ReadAll(got["file"])
		AssertNoError(t, err)
		AssertEqual(t, "Payload File", string(file))

		// Test _payload content
		payload, err := io.ReadAll(got["_payload"])
		AssertNoError(t, err)

		// Parse the JSON payload to verify the fields
		var parsedPayload struct {
			Alt     string `json:"alt"`
			Caption string `json:"caption"`
		}
		err = json.Unmarshal(payload, &parsedPayload)
		AssertNoError(t, err)

		AssertEqual(t, "John Doe", parsedPayload.Alt)
		AssertEqual(t, "Hello World", parsedPayload.Caption)
	})
}

type mockErrWriter struct{}

func (m mockErrWriter) Write(_ []byte) (n int, err error) {
	return 0, errors.New("mock write error")
}

func TestHandleFileUpload(t *testing.T) {
	t.Parallel()

	t.Run("File Closed", func(t *testing.T) {
		f, teardown, err := createTestFile(t, []byte("Payload File"))
		defer teardown()
		AssertNoError(t, err)

		AssertNoError(t, f.Close())
		err = handleFileUpload(nil, "file", f, MediaOptions{})
		AssertError(t, err)
	})

	t.Run("Test", func(t *testing.T) {
		t.Parallel()

		f, teardown, err := createTestFile(t, []byte("Payload File"))
		defer teardown()
		AssertNoError(t, err)

		w := multipart.NewWriter(&mockErrWriter{})
		err = handleFileUpload(w, "file", f, MediaOptions{})
		AssertError(t, err)
	})

	t.Run("Correct Override", func(t *testing.T) {
		t.Parallel()

		f, teardown, err := createTestFile(t, []byte("Payload File"))
		defer teardown()
		AssertNoError(t, err)

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		err = handleFileUpload(writer, "file.text", f, MediaOptions{
			FileNameOverride: "hello",
		})
		AssertNoError(t, err)
		AssertNoError(t, writer.Close())
		AssertEqual(t, true, strings.Contains(body.String(), "hello.txt"))
	})

	t.Run("MultiPart", func(t *testing.T) {
		t.Parallel()

		tempFile, teardown, err := createTestFile(t, []byte("Payload File"))
		defer teardown()
		AssertNoError(t, err)

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		err = handleFileUpload(writer, "file", tempFile, MediaOptions{})
		AssertNoError(t, err)
		AssertNoError(t, writer.Close())
		AssertEqual(t, true, strings.Contains(body.String(), "text/plain; charset=utf-8"))
	})
}

func TestFileNameFromURL(t *testing.T) {
	tt := map[string]struct {
		input string
		want  string
	}{
		"URL with filename": {
			input: "https://example.com/path/to/file.txt",
			want:  "file.txt",
		},
		"URL with trailing slash": {
			input: "https://example.com/path/to/",
			want:  "", // Expecting empty string as no filename is present
		},
		"URL with empty path": {
			input: "https://example.com/",
			want:  "", // Expecting empty string as no filename is present
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := fileNameFromURL(test.input)
			AssertEqual(t, test.want, got)
		})
	}
}
