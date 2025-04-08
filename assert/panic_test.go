package assert_test

import (
	"os"
	"os/exec"
	"testing"

	"github.com/fengdotdev/golibs-testing/assert"
)

func panicFn() {
	panic("panic")
}

func notPanicFn() {
	// do nothing
}

func TestPanic(t *testing.T) {
	assert.AssertPanic(t, panicFn)
}

func TestAssertPanicFailsWhenNotPanics(t *testing.T) {
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

func TestPanicWithMessage(t *testing.T) {
	assert.AssertPanicWithMessage(t, panicFn, "panic")
}

func TestAssertPanicWithMessageFailsWhenNotPanics(t *testing.T) {
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
