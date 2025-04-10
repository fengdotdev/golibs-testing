package assert

import "testing"

// Error checks if the error is not nil.
func Error(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Error("Expected an error, but got nil")
	}
}

func ErrorWithMessage(t *testing.T, err error, message string) {
	t.Helper()
	if err == nil {
		t.Error(message)
	}
}

func NoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
func NoErrorWithMessage(t *testing.T, err error, message string) {
	t.Helper()
	if err != nil {
		t.Error(message)
	}
}
