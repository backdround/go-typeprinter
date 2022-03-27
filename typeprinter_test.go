package typeprinter_test

import (
	"github.com/backdround/typeprinter"
	"testing"
)

func assertEqual(t *testing.T, expected interface{}, real interface{},
	failMessage string) {
	t.Helper()
	if expected != real {
		t.Errorf("%s\nexpected:%v\nreal:%v\n", failMessage, expected, real)
	}
}

func TestNil(t *testing.T) {
	assertEqual(t, "\n", typeprinter.Sprint(nil),
		"Nil representation is invalid")
}

func TestValue(t *testing.T) {
	assertEqual(t, "10\n", typeprinter.Sprint(10),
		"Value representation is invalid")
}

func TestString(t *testing.T) {
	assertEqual(t, "\"hi\"\n", typeprinter.Sprint("hi"),
		"String representation is invalid")
}

func TestLambdaEmptyStruct(t *testing.T) {
	representation := typeprinter.Sprint(struct{}{})
	assertEqual(t, representation, "{}\n", "Empty representation is invalid")
}

type emptyStruct struct{}

func TestUseTypeIfUnnamedStruct(t *testing.T) {
	representation := typeprinter.Sprint(emptyStruct{})
	assertEqual(t, "emptyStruct {}\n", representation,
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

	real := typeprinter.Sprint(p)

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
	assertEqual(t, expected, real, "Sturct representation is invalid")
}
