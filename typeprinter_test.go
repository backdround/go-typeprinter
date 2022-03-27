package typeprinter

import (
	"testing"
)

func TestNil(t *testing.T) {
	assertEqual(t, "\n", Sprint(nil),
		"Nil representation is invalid")
}

func TestValue(t *testing.T) {
	assertEqual(t, "10\n", Sprint(10),
		"Value representation is invalid")
}

func TestString(t *testing.T) {
	assertEqual(t, "\"hi\"\n", Sprint("hi"),
		"String representation is invalid")
}

func TestLambdaEmptyStruct(t *testing.T) {
	representation := Sprint(struct{}{})
	assertEqual(t, representation, "{}\n", "Empty representation is invalid")
}

type emptyStruct struct{}

func TestUseTypeIfUnnamedStruct(t *testing.T) {
	representation := Sprint(emptyStruct{})
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

	real := Sprint(p)

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
