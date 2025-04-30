package assert

import (
	"fmt"
	"testing"
)

func NotNil(t *testing.T, actual interface{}) {
	t.Helper()
	if actual == nil {
		t.Errorf("Expected not nil, but got nil")
	}
}

func NotNilWithMessage(t *testing.T, actual interface{}, message string, args ...any) {
	t.Helper()

	m := ""
	if len(args) > 0 {
		m = fmt.Sprintf(message, args...)
	} else {
		m = message
	}

	if actual == nil {
		t.Error(m)
	}
}
