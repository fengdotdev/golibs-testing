package assert_test

import (
	"testing"

	assert "github.com/fengdotdev/golibs-testing/assert"
)

func TestNil(t *testing.T) {
	assert.Nil(t, nil)
	assert.Nil(t, neverError)
	assert.NilWithMessage(t, nil, "This should be nil")
	assert.NilWithMessage(t, neverError, "This should be nil")
}

func TestNotNIl(t *testing.T) {
	assert.NotNil(t, errorSome)
	assert.NotNil(t, alwaysTrue)
	assert.NotNil(t, alwaysFalse)
	assert.NotNil(t, numberFive)
	assert.NotNil(t, numeroCinco)
	assert.NotNil(t, numberSix)
	assert.NotNil(t, numeroSeis)
	assert.NotNilWithMessage(t, errorSome, "This should not be nil")
	assert.NotNilWithMessage(t, alwaysTrue, "This should not be nil")
	assert.NotNilWithMessage(t, alwaysFalse, "This should not be nil")
	assert.NotNilWithMessage(t, numberFive, "This should not be nil")
	assert.NotNilWithMessage(t, numeroCinco, "This should not be nil")
	assert.NotNilWithMessage(t, numberSix, "This should not be nil")
	assert.NotNilWithMessage(t, numeroSeis, "This should not be nil")
	assert.NotNilWithMessage(t, t, "This should not be nil")
}
