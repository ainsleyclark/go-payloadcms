package payloadcms

import (
	"io"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestGetUploadValues(t *testing.T) {
	// Test case 1: Valid input with file and no struct
	f, err := os.Open("test.txt") // Replace with a valid file path
	if err != nil {
		t.Errorf("failed to open test file: %v", err)
	}
	defer f.Close()
	values, err := getUploadValues(f, nil)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	expectedValues := map[string]io.Reader{"file": f}
	if !reflect.DeepEqual(values, expectedValues) {
		t.Errorf("expected values: %v, got: %v", expectedValues, values)
	}

	// Test case 2: Missing file
	values, err = getUploadValues(nil, nil)
	if err == nil {
		t.Errorf("expected error for missing file")
	} else if err.Error() != "file is required" {
		t.Errorf("expected specific error message, got: %v", err)
	}

	// Test case 3: With a struct containing tagged fields (dummy struct)
	type MyStruct struct {
		Title string `json:"title"`
	}

	ms := MyStruct{Title: "Test Title"}
	values, err = getUploadValues(f, ms)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	expectedValues = map[string]io.Reader{
		"file":  f,
		"title": strings.NewReader("Test Title"),
	}
	if !reflect.DeepEqual(values, expectedValues) {
		t.Errorf("expected values: %v, got: %v", expectedValues, values)
	}
}
