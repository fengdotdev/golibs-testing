package assert_test

import (
	"testing"

	assert "github.com/fengdotdev/golibs-testing/assert"
)

func TestAssert_NotFail(t *testing.T) {

	assert.AssertNotFail(t, func(t *testing.T) {
		alwaysSucceed(t)
	})

	assert.AssertNotFail(t, func(t *testing.T) {

		neverPanicFn()
	})

}

func TestAssert_NotFailWithMessage(t *testing.T) {

	assert.AssertNotFailWithMessage(t, func(t *testing.T) {
		alwaysSucceed(t)
	}, "never fail")

	assert.AssertNotFailWithMessage(t, func(t *testing.T) {

		neverPanicFn()
	}, "never panic")

}
