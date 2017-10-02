package test

import (
"testing"
"golang.org/x/mobile/bind/objc/testpkg"
)

func TestHello(t *testing.T)  {
	expected := "Hello, Go!"
	actual := testpkg.Hello("Go")

	if actual != expected  {
		t.Errorf("Test failed, expected: '%s', got '%s'", expected, actual)
	}
}
