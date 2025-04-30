package assert

import (
	"reflect"
	"testing"
)

// equal checks if the expected and actual values are equal.

func Equal(t *testing.T, expected, actual interface{}) {
	t.Helper()

	eq := reflect.DeepEqual(expected, actual)

	if !eq {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func EqualWithMessage(t *testing.T, expected, actual interface{}, message string) {
	t.Helper()

	eq := reflect.DeepEqual(expected, actual)

	if !eq {
		t.Errorf(message)
	}
}

func NotEqual(t *testing.T, expected, actual interface{}) {
	t.Helper()

	eq := reflect.DeepEqual(expected, actual)
	if eq {
		t.Errorf("Expected %v to not be equal to %v", expected, actual)
	}
}

func NotEqualWithMessage(t *testing.T, expected, actual interface{}, message string) {
	t.Helper()
	eq := reflect.DeepEqual(expected, actual)
	if eq {
		t.Errorf(message)
	}
}
