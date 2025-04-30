package assert

import (
	"fmt"
	"testing"
)

// deprecated: NoError is deprecated, use NotError instead.
func NoError(t *testing.T, err error) {
	t.Helper()
	NotError(t, err)
}

// NotError asserts that the error is nil.
// If the error is not nil, it reports an error with the message "Expected no error, but got <error>".
func NotError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}



// deprecated: NoErrorWithMessage is deprecated, use NotErrorWithMessage instead.
func NoErrorWithMessage(t *testing.T, err error, message string) {
	t.Helper()
	NotErrorWithMessage(t, err, message)
}

// NotErrorWithMessage asserts that the error is nil, and if not, it reports an error with a custom message.
func NotErrorWithMessage(t *testing.T, err error, message string, args ...any) {
	t.Helper()
	m := ""
	if len(args) > 0 {
		m = fmt.Sprintf(message, args...)
	} else {
		m = message
	}

	if err != nil {
		t.Error(m)
	}
}
