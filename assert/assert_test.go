package assert_test

import (
	"errors"
	"testing"

	assert "github.com/fengdotdev/golibs-testing/assert"
)


//data


var numberFive = 5
var numeroCinco = 5
var numberSix = 6
var numeroSeis = 6
var alwaysTrue = true
var alwaysFalse = false
var neverError error = nil
var alwaysError = errors.New("error")






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

	assert.Error(t, alwaysError)

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
	assert.ErrorWithMessage(t, alwaysError, "This should be an error")
	assert.NilWithMessage(t, nil, "This should be nil")
	assert.NilWithMessage(t, neverError, "This should be nil")
	assert.FalseWithMessage(t, alwaysFalse, "This should be false")
	assert.TrueWithMessage(t, alwaysTrue, "This should be true")
	
}


