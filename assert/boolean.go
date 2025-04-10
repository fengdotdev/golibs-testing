package assert

import "testing"

func False(t *testing.T, actual bool) {
	t.Helper()
	if actual {
		t.Error("Expected false, but got true")
	}
}

func FalseWithMessage(t *testing.T, actual bool, message string) {
	t.Helper()
	if actual {
		t.Error(message)
	}
}

func True(t *testing.T, actual bool) {
	t.Helper()
	if !actual {
		t.Error("Expected true, but got false")
	}
}

func TrueWithMessage(t *testing.T, actual bool, message string) {
	t.Helper()
	if !actual {
		t.Error(message)
	}
}
