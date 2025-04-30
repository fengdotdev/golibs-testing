package assert_test

import (
	"testing"

	assert "github.com/fengdotdev/golibs-testing/assert"
)

func Test_Err(t *testing.T) {
	assert.Error(t, errorSome)
}

func Test_Err_withMsg(t *testing.T) {
	assert.ErrorWithMessage(t, errorSome, "This should not be nil")
}

