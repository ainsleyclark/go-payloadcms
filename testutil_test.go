package payloadcms

import "testing"

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func AssertNoError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}
