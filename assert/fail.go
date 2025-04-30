package assert

import (
	"fmt"
	"sync"
	"testing"
)

// deprecated: use AssertFail instead
func AssertFail(t *testing.T, innerTest func(t *testing.T)) {
	t.Helper()
	Fail(t, innerTest)
}

// Fail asserts that the innerTest function fails. It runs the innerTest in a goroutine and checks if it fails.
// If innerTest does not fail, it reports an error.
func Fail(t *testing.T, innerTest func(t *testing.T)) {
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
	if !failed {
		t.Error("expected innerTest to fail, but it passed")
	}
}

// deprecated: use AssertFailWithMessage instead
func AssertFailWithMessage(t *testing.T, innerTest func(t *testing.T), message string) {
	t.Helper()
	FailWithMessage(t, innerTest, message)
}

// FailWithMessage asserts that the innerTest function fails. It runs the innerTest in a goroutine and checks if it fails.
// If innerTest does not fail, it reports an error with the provided message.
func FailWithMessage(t *testing.T, innerTest func(t *testing.T), message string, args ...any) {
	t.Helper()

	m := ""
	if len(args) > 0 {
		m = fmt.Sprintf(message, args...)
	} else {
		m = message
	}

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
	if !failed {
		t.Error(m)
	}
}
