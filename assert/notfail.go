package assert

import (
	"fmt"
	"sync"
	"testing"
)

// deprecated: use NotFail instead
func AssertNotFail(t *testing.T, innerTest func(t *testing.T)) {
	t.Helper()
	NotFail(t, innerTest)
}

// NotFail assert a not fail function
// executes the function that u not expect to fail if it does fail the test fails
func NotFail(t *testing.T, innerTest func(t *testing.T)) {
	t.Helper()

	// Create a subtest to capture the failure state of innerTest
	failed := false
	t.Run("inner", func(t *testing.T) {

		wg := sync.WaitGroup{}
		wg.Add(1)

		go func() {
			t1 := &testing.T{}
			defer func() {
				wg.Done()
				if r := recover(); r != nil {
					// If innerTest panics, we consider it a failure
					failed = true
				} else {
					failed = t1.Failed()
				}

			}()
			// Run the inner test

			innerTest(t1)

		}()
		wg.Wait()

	})

	// If innerTest did not fail, report an error
	if failed {
		t.Error("expected innerTest to pass, but it failed")
	}
}

// deprecated: use NotFailWithMessage instead
func AssertNotFailWithMessage(t *testing.T, innerTest func(t *testing.T), message string) {
	t.Helper()
	NotFailWithMessage(t, innerTest, message)
}

// NotFailWithMessage assert a not fail function with message
// executes the function that u not expect to fail with message if it does fail the test fails
func NotFailWithMessage(t *testing.T, innerTest func(t *testing.T), message string, args ...any) {
	t.Helper()

	// Create a subtest to capture the failure state of innerTest
	failed := false
	t.Run("inner", func(t *testing.T) {

		wg := sync.WaitGroup{}
		wg.Add(1)

		go func() {
			t1 := &testing.T{}
			defer func() {
				wg.Done()
				if r := recover(); r != nil {
					// If innerTest panics, we consider it a failure
					failed = true
				} else {
					failed = t1.Failed()
				}

			}()
			// Run the inner test

			innerTest(t1)

		}()
		wg.Wait()

	})

	m := ""
	if len(args) > 0 {
		m = fmt.Sprintf(message, args...)
	} else {
		m = message
	}

	// If innerTest did not fail, report an error
	if failed {
		t.Error(m)
	}
}
