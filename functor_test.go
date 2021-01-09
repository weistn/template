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

func TestFunctor(t *testing.T) {
	funcMap := FuncMap{
		"threeAdic": threeAdic,
	}
	// template, err := New("root").Delims("{{", "}}").Funcs(funcMap).Parse("Hello {{ $one := call threeAdic 2 }} {{ $two := call $one 3 }} {{ printf \"Depp\" }}")
	template, err := New("root").Delims("{{", "}}").Funcs(funcMap).Parse("{{ $one := call . 2 }}{{ $two := call $one 3 }}{{ call $two 4 }}")
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

func TestFunctor2(t *testing.T) {
	funcMap := FuncMap{
		"threeAdic": threeAdic,
		"fn":        withFunctor,
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
