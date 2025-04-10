package assert_test

import (
	"testing"

	assert "github.com/fengdotdev/golibs-testing/assert"
)

func TestAssert_Fail(t *testing.T) {

	assert.AssertFail(t, func(t *testing.T) {
		alwaysFail(t)
	})

	assert.AssertFail(t, func(t *testing.T) {
		alwaysFail2(t)
	})

	assert.AssertFail(t, func(t *testing.T) {
		alwaysPanicFn()
	})

}

func TestAssert_FailWithMessage(t *testing.T) {

}
