package assert_test

import (
	"testing"

	assert "github.com/fengdotdev/golibs-testing/assert"
)



func TestAssert_Fail(t *testing.T) {

	assert.AssertFail(t, func(t *testing.T) {
		alwaysFail(t)
	})
	
}

func TestAssert_FailWithMessage(t *testing.T) {
	assert.AssertFailWithMessage(t, func(t *testing.T) {
		alwaysFail(t)
	}, "expected failure")

}



