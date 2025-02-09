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
