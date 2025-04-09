package assert_test

import (
	"os"
	"os/exec"
	"testing"

	"github.com/fengdotdev/golibs-testing/assert"
)



func Test_AssertPanic(t *testing.T) {
	assert.AssertPanic(t, panicFn)
}

func Test_AssertPanic_FailsWhenNotPanics(t *testing.T) {
	if os.Getenv("TEST_SHOULD_FAIL") == "1" {
		assert.AssertPanic(t, notPanicFn)
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestAssertPanicFailsWhenNotPanics")
	cmd.Env = append(os.Environ(), "TEST_SHOULD_FAIL=1")
	err := cmd.Run()

	if err == nil {
		t.Errorf("expected test to fail, but it passed")
	}
}

func Test_AssertPanicWithMessage(t *testing.T) {
	assert.AssertPanicWithMessage(t, panicFn, "panic")
}

func Test_AssertPanicWithMessageFails_WhenNotPanics(t *testing.T) {
	if os.Getenv("TEST_SHOULD_FAIL") == "1" {
		assert.AssertPanicWithMessage(t, notPanicFn, "panic")
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestAssertPanicWithMessageFailsWhenNotPanics")
	cmd.Env = append(os.Environ(), "TEST_SHOULD_FAIL=1")
	err := cmd.Run()

	if err == nil {
		t.Errorf("expected test to fail, but it passed")
	}
}
