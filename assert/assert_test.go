package assert_test

import (
	"errors"
	"testing"

	assert "github.com/fengdotdev/golibs-testing/assert"
)

func TestAssert(t *testing.T) {

	numberFive := 5
	numeroCinco := 5
	numberSix := 6
	numeroSeis := 6
	alwaysTrue := true
	alwaysFalse := false
	var neverError error = nil
	alwaysError := errors.New("error")

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

}
