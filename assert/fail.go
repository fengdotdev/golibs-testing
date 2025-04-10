package assert

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

// AssertShouldFail executes the innerTest function and expects it to fail.
func AssertFail(t *testing.T, innerTest func(t *testing.T)) {
	
	t.Helper()

	testName := "failCheck" + t.Name()

	if os.Getenv("FAIL_CHECK") == testName {
		innerTest(t)
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run="+testName)
	cmd.Env = append(os.Environ(), "FAIL_CHECK="+testName)
	output, err := cmd.CombinedOutput()

	if err == nil {
		t.Errorf("expected failure but test passed")
	}

	if !strings.Contains(string(output), "FAIL") {
		t.Errorf("expected output to contain 'FAIL', got: %s", string(output))
	}
}

func AssertFailWithMessage(t *testing.T, innerTest func(t *testing.T), message string) {
	t.Helper()

	testName := "failCheck" + t.Name()

	if os.Getenv("FAIL_CHECK") == testName {
		innerTest(t)
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run="+testName)
	cmd.Env = append(os.Environ(), "FAIL_CHECK="+testName)
	output, err := cmd.CombinedOutput()

	if err == nil {
		t.Errorf("expected failure but test passed")
	}

	if !strings.Contains(string(output), message) {
		t.Errorf("expected output to contain '%s', got: %s", message, string(output))
	}
}

func AssertNotFail(t *testing.T, innerTest func(t *testing.T)) {
	t.Helper()

	testName := "failCheck" + t.Name()

	if os.Getenv("FAIL_CHECK") == testName {
		innerTest(t)
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run="+testName)
	cmd.Env = append(os.Environ(), "FAIL_CHECK="+testName)
	output, err := cmd.CombinedOutput()

	if err != nil {
		t.Errorf("expected success but test failed: %s", string(output))
	}
}

func AssertNotFailWithMessage(t *testing.T, innerTest func(t *testing.T), message string) {
	t.Helper()

	testName := "failCheck" + t.Name()

	if os.Getenv("FAIL_CHECK") == testName {
		innerTest(t)
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run="+testName)
	cmd.Env = append(os.Environ(), "FAIL_CHECK="+testName)
	output, err := cmd.CombinedOutput()

	if err != nil {
		t.Errorf("expected success but test failed: %s", string(output))
	}

	if strings.Contains(string(output), message) {
		t.Errorf("expected output to not contain '%s', got: %s", message, string(output))
	}
}
