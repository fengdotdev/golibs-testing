package assert_test

import (
	"testing"

	assert "github.com/fengdotdev/golibs-testing/assert"
)

func Test_False(t *testing.T) {
	assert.False(t, alwaysFalse)
}

func Test_True(t *testing.T) {
	assert.True(t, alwaysTrue)
}

func Test_FalseWithMessage(t *testing.T) {
	assert.FalseWithMessage(t, alwaysFalse, "This should be false")
}


func Test_TrueWithMessage(t *testing.T) {
	assert.TrueWithMessage(t, alwaysTrue, "This should be true")
}
