package assert_test

import (
	"testing"

	assert "github.com/fengdotdev/golibs-testing/assert"
)

func Test_Equal(t *testing.T) {
	assert.Equal(t, 1, 1)
	assert.Equal(t, "a", "a")
	assert.Equal(t, true, true)
}

func Test_NotEqual(t *testing.T) {
	assert.NotEqual(t, 1, 2)
	assert.NotEqual(t, "a", "b")
	assert.NotEqual(t, true, false)
}

func Test_EqualWithMessage(t *testing.T) {
	assert.EqualWithMessage(t, 1, 1, "This should be equal")
	assert.EqualWithMessage(t, "a", "a", "This should be equal")
	assert.EqualWithMessage(t, true, true, "This should be equal")
}

func Test_NotEqualWithMessage(t *testing.T) {
	assert.NotEqualWithMessage(t, 1, 2, "This should not be equal")
	assert.NotEqualWithMessage(t, "a", "b", "This should not be equal")
	assert.NotEqualWithMessage(t, true, false, "This should not be equal")
}
