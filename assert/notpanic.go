package assert

import "testing"

// executes the function that u not expect to panic if it does panic the test fails
func AssertNotpanic(t *testing.T, neverPanicFn func()) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Expected not panic, but got panic")
		}
	}()
	neverPanicFn()
}

// executes the function that u not expect to panic with message if it does panic the test fails
func AssertNotpanicWithMessage(t *testing.T, neverPanicFn func(), message string) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf(message)
		}
	}()
	neverPanicFn()
}
