package assert_test

import (
	"os"
	"os/exec"
	"testing"

	"github.com/fengdotdev/golibs-testing/assert"
)

func TestNotPanic(t *testing.T) {
	assert.AssertNotPanic(t, neverPanicFn)

}

func TestAssert_NotPanic_FailsWhenPanics(t *testing.T) {
	if os.Getenv("TEST_SHOULD_FAIL") == "1" {
		assert.AssertNotPanic(t, alwaysPanicFn)
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestAssert_NotPanic_FailsWhenPanics")
	cmd.Env = append(os.Environ(), "TEST_SHOULD_FAIL=1")
	err := cmd.Run()

	if err == nil {
		t.Errorf("expected test to fail, but it passed")
	}
}

func TestNotPanicWithMessage(t *testing.T) {
	assert.AssertNotPanicWithMessage(t, neverPanicFn, "this should not panic, but it did")
}

func TestAssert_NotPanicWithMessage_FailsWhenPanics(t *testing.T) {
	if os.Getenv("TEST_SHOULD_FAIL") == "1" {
		assert.AssertNotPanicWithMessage(t, alwaysPanicFn, "this test must fail beause it panics , in a not panic assert")
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestAssert_NotPanicWithMessage_FailsWhenPanics")
	cmd.Env = append(os.Environ(), "TEST_SHOULD_FAIL=1")
	err := cmd.Run()

	if err == nil {
		t.Errorf("expected test to fail, but it passed")
	}
}
