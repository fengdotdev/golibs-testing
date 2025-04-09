package assert_test

import (
	"os"
	"os/exec"
	"testing"

	"github.com/fengdotdev/golibs-testing/assert"
)

func TestNotPanic(t *testing.T) {
	assert.AssertNotpanic(t, notPanicFn)

}

func TestAssertNotpanicFailsWhenPanics(t *testing.T) {
	if os.Getenv("TEST_SHOULD_FAIL") == "1" {
		assert.AssertNotpanic(t, panicFn)
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestAssertNotpanicFailsWhenPanics")
	cmd.Env = append(os.Environ(), "TEST_SHOULD_FAIL=1")
	err := cmd.Run()

	if err == nil {
		t.Errorf("expected test to fail, but it passed")
	}
}

func TestNotPanicWithMessage(t *testing.T) {
	assert.AssertNotpanicWithMessage(t, notPanicFn, "not panic")
}

func TestAssertNotpanicWithMessageFailsWhenPanics(t *testing.T) {
	if os.Getenv("TEST_SHOULD_FAIL") == "1" {
		assert.AssertNotpanicWithMessage(t, panicFn, "not panic")
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestAssertNotpanicWithMessageFailsWhenPanics")
	cmd.Env = append(os.Environ(), "TEST_SHOULD_FAIL=1")
	err := cmd.Run()

	if err == nil {
		t.Errorf("expected test to fail, but it passed")
	}
}
