package assert

import (
	"testing"
)


// Type asserts that the actual value is of the expected type.
// If the assertion fails, it reports an error with the message "Expected type <expectedType>, but got <actualType>".
func Type[T any](t *testing.T, actual interface{}, expectedType T) {
	t.Helper()

	s, ok := actual.(T)
	if !ok {
		t.Errorf("Expected type %T, but got %T", expectedType, s)
	} else {
	}
}

// TypeWithMessage asserts that the actual value is of the expected type.
// If the assertion fails, it reports an error with a custom message.
// The message can be formatted with additional arguments.
func TypeWithMessage[T any](t *testing.T, actual interface{}, expectedType T, message string, args ...any) {
	t.Helper()

	m := ""
	if len(args) > 0 {
		m = message
	} else {
		m = message
	}

	s, ok := actual.(T)
	if !ok {
		t.Errorf(m)
	} else {
		t.Logf("Type assertion successful: %T", s)
	}
}