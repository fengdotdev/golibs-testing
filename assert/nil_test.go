package assert_test

import (
	"testing"

	"github.com/fengdotdev/golibs-testing/assert"
)

func TestNil(t *testing.T) {
	assert.Nil(t, nil)
	assert.Nil(t, neverError)
	assert.NilWithMessage(t, nil, "This should be nil")
	assert.NilWithMessage(t, neverError, "This should be nil")
}
