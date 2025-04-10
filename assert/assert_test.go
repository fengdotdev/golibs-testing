package assert_test

import (
	"testing"

	assert "github.com/fengdotdev/golibs-testing/assert"
)

//data

func TestAssert(t *testing.T) {

	//tests
	assert.Equal(t, numberFive, numeroCinco)
	assert.Equal(t, numberSix, numeroSeis)
	assert.NotEqual(t, numberFive, numberSix)
	assert.NotEqual(t, numeroCinco, numeroSeis)

	assert.NotEqual(t, numberFive, numberSix)
	assert.NotEqual(t, alwaysTrue, alwaysFalse)

	assert.NoError(t, nil)
	assert.NoError(t, neverError)

	assert.Error(t, errorSome)

	assert.Nil(t, nil)
	assert.Nil(t, neverError)

	assert.False(t, alwaysFalse)
	assert.True(t, alwaysTrue)

}

func TestAssertWithMessage(t *testing.T) {

	assert.EqualWithMessage(t, numberFive, numeroCinco, "This should be equal")
	assert.EqualWithMessage(t, numberSix, numeroSeis, "This should be equal")
	assert.NotEqualWithMessage(t, numberFive, numberSix, "This should not be equal")
	assert.NotEqualWithMessage(t, numeroCinco, numeroSeis, "This should not be equal")
	assert.NotEqualWithMessage(t, numberFive, numberSix, "This should not be equal")
	assert.NotEqualWithMessage(t, alwaysTrue, alwaysFalse, "This should not be equal")
	assert.NoErrorWithMessage(t, nil, "This should not be an error")
	assert.NoErrorWithMessage(t, neverError, "This should not be an error")
	assert.ErrorWithMessage(t, errorSome, "This should be an error")
	assert.NilWithMessage(t, nil, "This should be nil")
	assert.NilWithMessage(t, neverError, "This should be nil")
	assert.FalseWithMessage(t, alwaysFalse, "This should be false")
	assert.TrueWithMessage(t, alwaysTrue, "This should be true")

}
