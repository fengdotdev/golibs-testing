package assert_test

import (
	"os"
	"os/exec"
	"testing"

	"github.com/fengdotdev/golibs-testing/assert"
)

func Test_AssertPanic(t *testing.T) {
	assert.AssertPanic(t, alwaysPanicFn)
}

func Test_AssertPanic_FailsWhenNotPanics(t *testing.T) {
	if os.Getenv("TEST_SHOULD_FAIL") == "1" {
		assert.AssertPanic(t, neverPanicFn)
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=Test_AssertPanic_FailsWhenNotPanics")
	cmd.Env = append(os.Environ(), "TEST_SHOULD_FAIL=1")
	err := cmd.Run()

	if err == nil {
		t.Errorf("expected test to fail, but it passed")
	}
}

func Test_AssertPanicWithMessage(t *testing.T) {
	assert.AssertPanicWithMessage(t, alwaysPanicFn, "this should panic, but it did not")
}

func Test_AssertPanicWithMessageFails_WhenNotPanics(t *testing.T) {
	if os.Getenv("TEST_SHOULD_FAIL") == "1" {
		assert.AssertPanicWithMessage(t, neverPanicFn, "this test must fail, because never panics expecting a panic")
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=Test_AssertPanicWithMessageFails_WhenNotPanics")
	cmd.Env = append(os.Environ(), "TEST_SHOULD_FAIL=1")
	err := cmd.Run()

	if err == nil {
		t.Errorf("expected test to fail, but it passed")
	}
}
