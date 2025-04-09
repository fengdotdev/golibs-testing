package assert_test

import (
	"testing"

	assert "github.com/fengdotdev/golibs-testing/assert"
)

func alwaysFail(t *testing.T) {
	t.Errorf("expected failure")
}
func alwaysSucceed(t *testing.T) {
	t.Log("expected success")
}

func TestAssertShouldFail(t *testing.T) {

	assert.AssertShouldFail(t, func(t *testing.T) {
		alwaysFail(t)
	})
	
}

func TestAssertShouldFailWithMessage(t *testing.T) {
	assert.AssertShouldFailWithMessage(t, func(t *testing.T) {
		alwaysFail(t)
	}, "expected failure")

}



