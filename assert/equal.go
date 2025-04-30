package assert

import (
	"fmt"
	"reflect"
	"testing"
)

// equal asserts that two values are equal.
// If they are not equal, it reports an error with the message "Expected <expected>, but got <actual>".
func Equal(t *testing.T, expected, actual interface{}) {
	t.Helper()

	eq := reflect.DeepEqual(expected, actual)

	if !eq {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}


// EqualWithMessage asserts that two values are equal, and if not, it reports an error with a custom message.
// The message can be formatted with additional arguments if provided.
func EqualWithMessage(t *testing.T, expected, actual interface{}, message string, args ...any) {
	t.Helper()

	m := ""
	if len(args) > 0 {
		m = fmt.Sprintf(message, args...)
	} else {
		m = message
	}

	eq := reflect.DeepEqual(expected, actual)

	if !eq {
		t.Errorf(m)
	}
}

