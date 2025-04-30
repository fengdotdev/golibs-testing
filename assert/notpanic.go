package assert

import (
	"fmt"
	"testing"
)

// Deprecated: use NotPanic instead
func AssertNotPanic(t *testing.T, neverPanicFn func()) {
	t.Helper()
	NotPanic(t, neverPanicFn)
}

// NotPanic assert a not panic function
// executes the function that u not expect to panic if it does panic the test fails
func NotPanic(t *testing.T, neverPanicFn func()) {
	t.Helper()
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Expected not panic, but got panic")
		}
	}()
	neverPanicFn()
}

// Deprecated: use NotPanicWithMessage instead
func AssertNotPanicWithMessage(t *testing.T, neverPanicFn func(), message string) {
	t.Helper()
	NotPanicWithMessage(t, neverPanicFn, message)
}

// NotPanicWithMessage assert a not panic function with message
// executes the function that u not expect to panic with message if it does panic the test fails
func NotPanicWithMessage(t *testing.T, neverPanicFn func(), message string, args ...any) {
	t.Helper()

	m := ""
	if len(args) > 0 {
		m = fmt.Sprintf(message, args...)
	} else {
		m = message
	}

	defer func() {
		if r := recover(); r != nil {
			t.Error(m)
		}
	}()
	neverPanicFn()
}
