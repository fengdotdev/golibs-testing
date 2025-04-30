package assert

import (
	"fmt"
	"testing"
)



// False asserts that the actual boolean value is false, and if not, it reports an error.
// The error message is "Expected false, but got true" if the assertion fails.
func False(t *testing.T, actual bool) {
	t.Helper()
	if actual {
		t.Error("Expected false, but got true")
	}
}


// FalseWithMessage asserts that the actual boolean value is false, and if not, it reports an error with a custom message.
// The message can be formatted with additional arguments if provided.
func FalseWithMessage(t *testing.T, actual bool, message string, args ...any) {
	t.Helper()

	m := ""
	if len(args) > 0 {
		m = fmt.Sprintf(message, args...)
	} else {
		m = message
	}

	if actual {
		t.Error(m)
	}
}



// True asserts that the actual boolean value is true, and if not, it reports an error.
func True(t *testing.T, actual bool) {
	t.Helper()
	if !actual {
		t.Error("Expected true, but got false")
	}
}


// TrueWithMessage asserts that the actual boolean value is true, and if not, it reports an error with a custom message.
// The message can be formatted with additional arguments if provided.
func TrueWithMessage(t *testing.T, actual bool, message string, args ...any) {
	t.Helper()

	m := ""
	if len(args) > 0 {
		m = fmt.Sprintf(message, args...)
	} else {
		m = message
	}

	if !actual {
		t.Error(m)
	}
}
