package assert

import "testing"

// Nil checks if the actual value is nil DUH
func Nil(t *testing.T, actual interface{}) {
	t.Helper()
	if actual != nil {
		t.Errorf("Expected nil, but got %v", actual)
	}
}

func NilWithMessage(t *testing.T, actual interface{}, message string) {
	t.Helper()
	if actual != nil {
		t.Error(message)
	}
}

func NotNil(t *testing.T, actual interface{}) {
	t.Helper()
	if actual == nil {
		t.Errorf("Expected not nil, but got nil")
	}
}

func NotNilWithMessage(t *testing.T, actual interface{}, message string) {
	t.Helper()
	if actual == nil {
		t.Error(message)
	}
}
