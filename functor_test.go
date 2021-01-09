package template

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

func threeAdic(a int, b int, c int) int {
	println("Hello from threeAdic")
	return a + b + c
}

func withFunctor(f Functor) (reflect.Value, error) {
	if f == nil {
		return reflect.Value{}, fmt.Errorf("No functor")
	}
	println("Hello from withFunctor")
	return f.Call(reflect.ValueOf(1), reflect.ValueOf(2), reflect.ValueOf(3))
}

func prefix(count int, str string) string {
	return str[:count]
}

func TestCallFunctor(t *testing.T) {
	template, err := New("root").Delims("{{", "}}").Parse("{{ $one := call . 2 }}{{ $two := call $one 3 }}{{ call $two 4 }}")
	if err != nil {
		t.Fatal(err)
	}
	w := bytes.NewBuffer(nil)
	err = template.Execute(w, threeAdic)
	if err != nil {
		t.Fatal(err)
	}
	if w.String() != "9" {
		println(">>>>" + w.String() + "<<<<<")
		t.Fatal(w.String())
	}
}

// Test that a Go function can be passed as argument to another function, which expects a Functor instead of a
// go function as argument. The template engine will automatically wrap the Go function inside a Functor.
func TestFunctorConversion(t *testing.T) {
	funcMap := FuncMap{
		"fn": withFunctor,
	}
	// template, err := New("root").Delims("{{", "}}").Funcs(funcMap).Parse("Hello {{ $one := call threeAdic 2 }} {{ $two := call $one 3 }} {{ printf \"Depp\" }}")
	template, err := New("root").Delims("{{", "}}").Funcs(funcMap).Parse("{{ fn . }}")
	if err != nil {
		t.Fatal(err)
	}
	w := bytes.NewBuffer(nil)
	err = template.Execute(w, threeAdic)
	if err != nil {
		t.Fatal(err)
	}
	if w.String() != "6" {
		println(">>>>" + w.String() + "<<<<<")
		t.Fatal(w.String())
	}
}

func TestEvalFunctor(t *testing.T) {
	funcMap := FuncMap{
		"threeAdic": threeAdic,
	}
	template, err := New("root").Delims("{{", "}}").Funcs(funcMap).Parse("{{ $one := threeAdic 2 }}{{ $two := call $one 3 }}{{ call $two 4 }}")
	if err != nil {
		t.Fatal(err)
	}
	w := bytes.NewBuffer(nil)
	err = template.Execute(w, threeAdic)
	if err != nil {
		t.Fatal(err)
	}
	if w.String() != "9" {
		println(">>>>" + w.String() + "<<<<<")
		t.Fatal(w.String())
	}
}

func TestMapFunc(t *testing.T) {
	funcMap := FuncMap{
		"prefix": prefix,
	}
	template, err := New("root").Delims("{{", "}}").Funcs(funcMap).Parse("{{ range map (prefix 1) . }}{{.}}{{end}}")
	if err != nil {
		t.Fatal(err)
	}
	w := bytes.NewBuffer(nil)
	err = template.Execute(w, []string{"Hello", "cruel", "world"})
	if err != nil {
		t.Fatal(err)
	}
	if w.String() != "Hcw" {
		println(">>>>" + w.String() + "<<<<<")
		t.Fatal(w.String())
	}
}
