package testutil

import "testing"

func AssertEqual(t *testing.T, expected, actual any, message string) {
	if expected != actual {
		t.Errorf("%s: expected %v, got %v", message, expected, actual)
	}
}
