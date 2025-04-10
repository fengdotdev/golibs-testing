package assert

import "testing"

// executes the function that u not expect to panic if it does panic the test fails
func AssertNotPanic(t *testing.T, neverPanicFn func()) {
	t.Helper()
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Expected not panic, but got panic")
		}
	}()
	neverPanicFn()
}

// executes the function that u not expect to panic with message if it does panic the test fails
func AssertNotPanicWithMessage(t *testing.T, neverPanicFn func(), message string) {
	t.Helper()
	defer func() {
		if r := recover(); r != nil {
			t.Error(message)
		}
	}()
	neverPanicFn()
}
