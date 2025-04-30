package assert

import (
	"fmt"
	"testing"
)

// Deprecated: use Panic instead
func AssertPanic(t *testing.T, alwaysPanicFn func()) {
	Panic(t, alwaysPanicFn)
}

// Panic assert a panic function
// executes the function that u expect to panic if it does not panic the test fails
func Panic(t *testing.T, alwaysPanicFn func()) {
	t.Helper()
	defer func() {
		if r := recover(); r != nil {
			t.Log("successfully panicked")
			return
		}
	}()
	alwaysPanicFn()
	t.Errorf("Expected panic, but did not panic")
}

// Deprecated: use PanicWithMessage instead
func AssertPanicWithMessage(t *testing.T, alwaysPanicFn func(), message string) {
	PanicWithMessage(t, alwaysPanicFn, message)
}

// PanicWithMessage assert a panic function with message
// executes the function that u expect to panic with message if it does not panic the test fails
func PanicWithMessage(t *testing.T, alwaysPanicFn func(), message string, args ...any) {

	t.Helper()
	defer func() {
		if r := recover(); r != nil {
			t.Log("successfully panicked")
			return
		}
	}()
	alwaysPanicFn()

	m := ""

	if len(args) > 0 {
		m = fmt.Sprintf(message, args...)
	} else {
		m = message
	}

	t.Error(m)
}
