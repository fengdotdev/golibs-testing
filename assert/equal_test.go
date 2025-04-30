package assert_test

import (
	"testing"

	assert "github.com/fengdotdev/golibs-testing/assert"
)

func Test_Equal(t *testing.T) {
	assert.Equal(t, 1, 1)
	assert.Equal(t, "a", "a")
	assert.Equal(t, true, true)
	assert.Equal(t, []int{1, 2}, []int{1, 2})
	assert.Equal(t, map[string]int{"a": 1}, map[string]int{"a": 1})
	assert.Equal(t, struct{ A int }{A: 1}, struct{ A int }{A: 1})
	assert.Equal(t, nil, nil)
	assert.Equal(t, []int{}, []int{})
	assert.Equal(t, map[string]int{}, map[string]int{})
	assert.Equal(t, struct{}{}, struct{}{})
}


func Test_EqualWithMessage(t *testing.T) {
	assert.EqualWithMessage(t, 1, 1, "This should be equal")
	assert.EqualWithMessage(t, "a", "a", "This should be equal")
	assert.EqualWithMessage(t, true, true, "This should be equal")
	assert.EqualWithMessage(t, []int{1, 2}, []int{1, 2}, "This should be equal")
	assert.EqualWithMessage(t, map[string]int{"a": 1}, map[string]int{"a": 1}, "This should be equal")
	assert.EqualWithMessage(t, struct{ A int }{A: 1}, struct{ A int }{A: 1}, "This should be equal")
	assert.EqualWithMessage(t, nil, nil, "This should be equal")
	assert.EqualWithMessage(t, []int{}, []int{}, "This should be equal")
	assert.EqualWithMessage(t, map[string]int{}, map[string]int{}, "This should be equal")
	assert.EqualWithMessage(t, struct{}{}, struct{}{}, "This should be equal")
}