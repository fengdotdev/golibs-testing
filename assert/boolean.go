package assert

import "testing"

func False(t *testing.T, actual bool) {
	if actual {
		t.Errorf("Expected false, but got true")
	}
}

func FalseWithMessage(t *testing.T, actual bool, message string) {
	if actual {
		t.Errorf(message)
	}
}

func True(t *testing.T, actual bool) {
	if !actual {
		t.Errorf("Expected true, but got false")
	}
}

func TrueWithMessage(t *testing.T, actual bool, message string) {
	if !actual {
		t.Errorf(message)
	}
}
