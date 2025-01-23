package assert

import "testing"

func Equal(t *testing.T, expected, actual interface{}) {
	if expected != actual {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func NotEqual(t *testing.T, expected, actual interface{}) {
	if expected == actual {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}

func Error(t *testing.T, err error) {
	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}
}

func NoError(t *testing.T, err error) {
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}

func Nil(t *testing.T, actual interface{}) {
	if actual != nil {
		t.Errorf("Expected nil, but got %v", actual)
	}
}
