package typeprinter

import (
	"testing"
)

func TestNil(t *testing.T) {
	assertEqual(t, "", Sprint(nil), "Nil representation is invalid")
}

func TestValue(t *testing.T) {
	assertEqual(t, "10", Sprint(10), "Value representation is invalid")
}

func TestString(t *testing.T) {
	assertEqual(t, `"hi"`, Sprint("hi"), "String representation is invalid")
}

func TestLambdaEmptyStruct(t *testing.T) {
	representation := Sprint(struct{}{})
	assertEqual(t, representation, "{}", "Empty representation is invalid")
}

type emptyStruct struct{}

func TestUseTypeIfUnnamedStruct(t *testing.T) {
	representation := Sprint(emptyStruct{})
	assertEqual(t, "emptyStruct {}", representation,
		"Unnamed sturct denotion is invalid")
}
