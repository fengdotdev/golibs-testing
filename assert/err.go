package assert

import (
	"fmt"
	"testing"
)

// Error asserts that the actual error is not nil, and if it is, it reports an error.
func Error(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Error("Expected an error, but got nil")
	}
}

// ErrorWithMessage asserts that the actual error is not nil, and if it is, it reports an error with a custom message.
// The message can be formatted with additional arguments if provided.
func ErrorWithMessage(t *testing.T, err error, message string, args ...any) {
	t.Helper()

	m := ""
	if len(args) > 0 {
		m = fmt.Sprintf(message, args...)
	} else {
		m = message
	}

	if err == nil {
		t.Error(m)
	}
}

