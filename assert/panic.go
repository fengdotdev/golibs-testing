package assert

import "testing"


// executes the function that u expect to panic if it does not panic the test fails
func AssertPanic(t *testing.T, alwaysPanicFn func()) {
	defer func() {
		if r := recover(); r != nil {
			t.Log("successfully panicked")
			return
		}
	}()
	alwaysPanicFn()
	t.Errorf("Expected panic, but did not panic")
}


// executes the function that u expect to panic with message if it does not panic the test fails
func AssertPanicWithMessage(t *testing.T, alwaysPanicFn func(), message string) {
	defer func() {
		if r := recover(); r != nil {
			t.Log("successfully panicked")
			return
		}
	}()
	alwaysPanicFn()
	t.Errorf(message)
}


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
