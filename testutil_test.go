package payloadcms

import "testing"

func AssertEqual[T comparable](t *testing.T, want, got T) {
	t.Helper()
	if got != want {
		t.Errorf("wanted %v, got %v", got, want)
	}
}

func AssertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func AssertError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Errorf("expected error, got %v", err)
	}
}
