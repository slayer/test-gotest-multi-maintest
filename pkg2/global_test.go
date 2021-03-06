package pkg2

import (
	"fmt"
	"os"
	"testing"
)

func TestFooFunc(t *testing.T) {
	if FooFunc() != "Foo" {
		t.Error("FooFunc fail")
	}
}

func TestBarFunc(t *testing.T) {
	if BarFunc() != "Bar" {
		t.Error("BarFunc fail")
	}
}

func TestMain(m *testing.M) {
	fmt.Println("pkg2 TestMain() setup")
	code := m.Run()
	fmt.Println("pkg2 TestMain() teardown")
	os.Exit(code)
}
