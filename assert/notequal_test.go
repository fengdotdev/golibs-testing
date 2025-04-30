package assert_test

import (
	"testing"

	assert "github.com/fengdotdev/golibs-testing/assert"
)

func Test_NotEqual(t *testing.T) {
	assert.NotEqual(t, 1, 2)
	assert.NotEqual(t, "a", "b")
	assert.NotEqual(t, true, false)
	assert.NotEqual(t, []int{1, 2}, []int{2, 1})
	assert.NotEqual(t, map[string]int{"a": 1}, map[string]int{"b": 1})
	assert.NotEqual(t, struct{ A int }{A: 1}, struct{ A int }{A: 2})
	assert.NotEqual(t, nil, 1)
	assert.NotEqual(t, []int{1}, nil)
	assert.NotEqual(t, map[string]int{"a": 1}, nil)
	assert.NotEqual(t, struct{ A int }{A: 1}, nil)
	assert.NotEqual(t, nil, struct{}{})
	assert.NotEqual(t, []int{1}, []int{})
	assert.NotEqual(t, map[string]int{"a": 1}, map[string]int{})
	assert.NotEqual(t, struct{ A int }{A: 1}, struct{}{})
}

func Test_NotEqualWithMessage(t *testing.T) {
	assert.NotEqualWithMessage(t, 1, 2, "This should not be equal")
	assert.NotEqualWithMessage(t, "a", "b", "This should not be equal")
	assert.NotEqualWithMessage(t, true, false, "This should not be equal")
	assert.NotEqualWithMessage(t, []int{1, 2}, []int{2, 1}, "This should not be equal")
	assert.NotEqualWithMessage(t, map[string]int{"a": 1}, map[string]int{"b": 1}, "This should not be equal")
	assert.NotEqualWithMessage(t, struct{ A int }{A: 1}, struct{ A int }{A: 2}, "This should not be equal")
	assert.NotEqualWithMessage(t, nil, 1, "This should not be equal")
}
