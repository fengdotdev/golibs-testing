package assert

import (
	"fmt"
	"reflect"
	"testing"
)

// NotEqual asserts that two values are not equal.
// If they are equal, it reports an error with the message "Expected <expected> to not be equal to <actual>".
func NotEqual(t *testing.T, expected, actual interface{}) {
	t.Helper()

	eq := reflect.DeepEqual(expected, actual)
	if eq {
		t.Errorf("Expected %v to not be equal to %v", expected, actual)
	}
}

// NotEqualWithMessage asserts that two values are not equal
// and if they are equal, it reports an error with a custom message.
// The message can be formatted with additional arguments if provided.
func NotEqualWithMessage(t *testing.T, expected, actual interface{}, message string, args ...any) {
	t.Helper()

	m := ""
	if len(args) > 0 {
		m = fmt.Sprintf(message, args...)
	} else {
		m = message
	}

	eq := reflect.DeepEqual(expected, actual)
	if eq {
		t.Errorf(m)
	}
}
