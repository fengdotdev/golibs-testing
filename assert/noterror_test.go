package assert_test

import (
	"testing"

	assert "github.com/fengdotdev/golibs-testing/assert"
)

func Test_Not_Err(t *testing.T) {
	assert.NoError(t, nil)
}

func Test_Not_Err_withMsg(t *testing.T) {
	assert.NoErrorWithMessage(t, nil, "This should be nil")
}
