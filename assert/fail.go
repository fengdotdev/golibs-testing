package assert

import (

	"sync"
	"testing"
)

// AssertFail executes innerTest and expects it to fail. If innerTest does not fail, the test fails.
func AssertFail(t *testing.T, innerTest func(t *testing.T)) {
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

func AssertFailWithMessage(t *testing.T, innerTest func(t *testing.T), message string) {
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
		t.Error(message)
	}
}
