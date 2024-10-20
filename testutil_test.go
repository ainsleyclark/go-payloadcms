package payloadcms

import (
	"strings"
	"testing"
)

// AssertEqual checks if two values are equal.
func AssertEqual[T comparable](t *testing.T, want, got T) {
	t.Helper()
	if got != want {
		t.Errorf("wanted %v, got %v", want, got)
	}
}

// AssertNoError checks if the error is nil.
func AssertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

// AssertError checks if an error is not nil.
func AssertError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

// AssertContains checks if the `substr` is contained in `str`.
func AssertContains(t *testing.T, str, substr string) {
	t.Helper()
	if !strings.Contains(str, substr) {
		t.Errorf("expected %q to contain %q", str, substr)
	}
}
