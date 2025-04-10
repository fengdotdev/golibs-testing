package assert_test

import (
	"errors"
	"testing"
)

// data for testing

var numberFive = 5
var numeroCinco = 5
var numberSix = 6
var numeroSeis = 6
var alwaysTrue = true
var alwaysFalse = false

// this is a nil error
var neverError error = nil

// this is a non-nil error
var errorSome error = errors.New("error")

// this function always panics
func alwaysPanicFn() {
	panic("panic")
}

// this function does not panic
func neverPanicFn() {
	// do nothing
}

// this function always fails
func alwaysFail(t *testing.T) {
	t.Error("expected failure")
}

// this function always succeeds
func alwaysSucceed(t *testing.T) {
	t.Log("expected success")
}
