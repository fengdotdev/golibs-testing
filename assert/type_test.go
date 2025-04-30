package assert_test

import (
	"testing"

	"github.com/fengdotdev/golibs-testing/assert"
)

func TestType(t *testing.T) {

	// ok

	assert.Type(t, 1, 1)
	assert.Type(t, "a", "a")
	assert.Type(t, true, true)
	assert.Type(t, []int{1, 2}, []int{1, 2})
	assert.Type(t, map[string]int{"a": 1}, map[string]int{"a": 1})
	assert.Type(t, struct{ A int }{A: 1}, struct{ A int }{A: 1})
	assert.Type(t, []int{}, []int{})
	assert.Type(t, []int{1, 2}, []int{1, 2, 3}) // same type, different length

	f1 := Foo{A: 1, B: "a", C: true}
	f2 := Foo{A: 1, B: "a", C: true}
	assert.Type(t, f1, f2) // same type,

	// not ok

	assert.AssertFailWithMessage(t, func(t *testing.T) {
		assert.Type(t, 1, "1")
	}, "Expected to fail, but it passed int != string")

	assert.AssertFailWithMessage(t, func(t *testing.T) {
		assert.Type(t, "a", 1)
	}, "Expected to fail, but it passed string != int")

	assert.AssertFailWithMessage(t, func(t *testing.T) {
		assert.Type(t, true, 1)
	}, "Expected to fail, but it passed bool != int")

	bar := Bar{A: 1, B: "a", C: true}
	foo := Foo{A: 1, B: "a", C: true}

	assert.AssertFailWithMessage(t, func(t *testing.T) {
		assert.Type(t, bar, foo)
	}, "Expected to fail, but it passed Bar != Foo")
}

func TestTypeWithMessage(t *testing.T) {

	assert.TypeWithMessage(t, 1, 1, "Expected pass, but it failed 1 == 1")

	assert.TypeWithMessage(t, "a", "a", "Expected pass, but it failed a == a")
	assert.TypeWithMessage(t, true, true, "Expected pass, but it failed true == true")

	assert.TypeWithMessage(t, []int{1, 2}, []int{1, 2}, "Expected pass, but it failed []int{1, 2} == []int{1, 2}")
	assert.TypeWithMessage(t, map[string]int{"a": 1}, map[string]int{"a": 1}, "Expected pass, but it failed map[string]int{\"a\": 1} == map[string]int{\"a\": 1}")

	assert.TypeWithMessage(t, struct{ A int }{A: 1}, struct{ A int }{A: 1}, "Expected pass, but it failed struct{ A int }{A: 1} == struct{ A int }{A: 1}")

	assert.TypeWithMessage(t, []int{}, []int{}, "Expected pass, but it failed []int{} == []int{}")

	assert.AssertFailWithMessage(t, func(t *testing.T) {
		assert.TypeWithMessage(t, 1, "1", "Expected to fail, but it passed int != string")
	}, "Expected to fail, but it passed int != string")

	assert.AssertFailWithMessage(t, func(t *testing.T) {
		assert.TypeWithMessage(t, "a", 1, "Expected to fail, but it passed string != int")
	}, "Expected to fail, but it passed string != int")

	assert.AssertFailWithMessage(t, func(t *testing.T) {
		assert.TypeWithMessage(t, true, 1, "Expected to fail, but it passed bool != int")
	}, "Expected to fail, but it passed bool != int")
}
