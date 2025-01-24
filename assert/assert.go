package assert

import "testing"

// equal checks if the expected and actual values are equal.

func Equal(t *testing.T, expected, actual interface{}) {
	if expected != actual {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func EqualWithMessage(t *testing.T, expected, actual interface{}, message string) {
	if expected != actual {
		t.Errorf(message)
	}
}

func NotEqual(t *testing.T, expected, actual interface{}) {
	if expected == actual {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func NotEqualWithMessage(t *testing.T, expected, actual interface{}, message string) {
	if expected == actual {
		t.Errorf(message)
	}
}

// Error checks if the error is not nil.
func Error(t *testing.T, err error) {
	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}
}

func ErrorWithMessage(t *testing.T, err error, message string) {
	if err == nil {
		t.Errorf(message)
	}
}

func NoError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
func NoErrorWithMessage(t *testing.T, err error, message string) {
	if err != nil {
		t.Errorf(message)
	}
}

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

func False(t *testing.T, actual bool) {
	if actual {
		t.Errorf("Expected false, but got true")
	}
}

func FalseWithMessage(t *testing.T, actual bool, message string) {
	if actual {
		t.Errorf(message)
	}
}

func True(t *testing.T, actual bool) {
	if !actual {
		t.Errorf("Expected true, but got false")
	}
}

func TrueWithMessage(t *testing.T, actual bool, message string) {
	if !actual {
		t.Errorf(message)
	}
}
