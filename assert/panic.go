package assert

import "testing"


// executes the function that u expect to panic if it does not panic the test fails
func AssertPanic(t *testing.T, alwaysPanicFn func()) {
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


// executes the function that u expect to panic with message if it does not panic the test fails
func AssertPanicWithMessage(t *testing.T, alwaysPanicFn func(), message string) {
	t.Helper()
	defer func() {
		if r := recover(); r != nil {
			t.Log("successfully panicked")
			return
		}
	}()
	alwaysPanicFn()
	t.Error(message)
}


