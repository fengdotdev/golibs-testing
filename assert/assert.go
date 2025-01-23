package assert

import "testing"

// equal checks if the expected and actual values are equal.

func Equal(t *testing.T, expected, actual interface{}) {
	if expected != actual {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func NotEqual(t *testing.T, expected, actual interface{}) {
	if expected == actual {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

// Error checks if the error is not nil.
func Error(t *testing.T, err error) {
	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}
}

func NoError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}

// Nil checks if the actual value is nil DUH
func Nil(t *testing.T, actual interface{}) {
	if actual != nil {
		t.Errorf("Expected nil, but got %v", actual)
	}
}

func False(t *testing.T, actual bool) {
	if actual {
		t.Errorf("Expected false, but got true")
	}
}

func True(t *testing.T, actual bool) {
	if !actual {
		t.Errorf("Expected true, but got false")
	}
}
