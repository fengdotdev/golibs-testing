package assert

import "testing"

// Nil checks if the actual value is nil DUH
func Nil(t *testing.T, actual interface{}) {
	if actual != nil {
		t.Errorf("Expected nil, but got %v", actual)
	}
}

func NilWithMessage(t *testing.T, actual interface{}, message string) {
	if actual != nil {
		t.Errorf(message)
	}
}

func NotNil(t *testing.T, actual interface{}) {
	if actual == nil {
		t.Errorf("Expected not nil, but got nil")
	}
}

func NotNilWithMessage(t *testing.T, actual interface{}, message string) {
	if actual == nil {
		t.Errorf(message)
	}
}
