package structprinter_test

import (
	"github.com/backdround/structprinter"
	"testing"
)

func assertEqual(t *testing.T, real interface{}, expected interface{},
	failMessage string) {
	t.Helper()
	if expected != real {
		t.Errorf("%s\nexpected:%v\nreal:%v\n", failMessage, expected, real)
	}
}

func TestNil(t *testing.T) {
	assertEqual(t, structprinter.Sprint(nil), "\n",
		"Nil representation is invalid")
}

func TestValue(t *testing.T) {
	assertEqual(t, structprinter.Sprint(10), "10\n",
		"Value representation is invalid")
}

func TestString(t *testing.T) {
	assertEqual(t, structprinter.Sprint("hi"), "\"hi\"\n",
		"String representation is invalid")
}

func TestLambdaEmptyStruct(t *testing.T) {
	representation := structprinter.Sprint(struct{}{})
	assertEqual(t, representation, "{}\n", "Empty representation is invalid")
}

type emptyStruct struct{}

func TestUseTypeIfUnnamedStruct(t *testing.T) {
	representation := structprinter.Sprint(emptyStruct{})
	assertEqual(t, representation, "emptyStruct {}\n",
		"Unnamed sturct denotion is invalid")
}

type person struct {
	name string
	age  int
	work struct {
		post  string
		floor int
	}
}

func TestStruct(t *testing.T) {
	p := person{
		name: "bob",
		age:  20,
		work: struct {
			post  string
			floor int
		}{
			post:  "boss",
			floor: 32,
		},
	}

	real := structprinter.Sprint(p)

	expected :=
		`person {
	name: "bob"
	age: 20
	work {
		post: "boss"
		floor: 32
	}
}
`
	assertEqual(t, real, expected, "Sturct representation is invalid")
}
