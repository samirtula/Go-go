package main

import (
	"testing"
)

// В Go принято располагать файлы с unit-тестами непосредственно в пакете, функции которого тестируются
func TestFooBarFunc(t *testing.T) {
	expectedFooResult := "bar"
	if actualFooResult := fooBar(); actualFooResult != expectedFooResult {
		t.Errorf("expected %s; got: %s", expectedFooResult, actualFooResult)
	}
}
