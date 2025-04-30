package assert

import (
	"fmt"
	"testing"
)

// Nil asserts that the actual value is nil.
// If the actual value is not nil, it reports an error with the message "Expected nil, but got <value>".
func Nil(t *testing.T, actual interface{}) {
	t.Helper()
	if actual != nil {
		t.Errorf("Expected nil, but got %v", actual)
	}
}

// NilWithMessage asserts that the actual value is nil, and if not, it reports an error with a custom message.
// The message can be formatted with additional arguments if provided.
func NilWithMessage(t *testing.T, actual interface{}, message string, args ...any) {
	t.Helper()

	m := ""
	if len(args) > 0 {
		m = fmt.Sprintf(message, args...)
	} else {
		m = message
	}

	if actual != nil {
		t.Error(m)
	}
}
