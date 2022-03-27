package typeprinter

import (
	"reflect"
	"testing"
)

type testStruct struct {
	name string
	age  int
}

var testValue = testStruct{"bob", 20}

func TestRootProperties(t *testing.T) {
	e := newElement(testValue, "bob-element")

	assertEqual(t, "bob-element", e.Name(), "Root name invalid")
	assertEqual(t, "{bob 20}", e.Value(), "Root value invalid")
	assertEqual(t, "testStruct", e.Type(), "Root type invald")
	assertEqual(t, reflect.Struct, e.Kind(), "Struct kind is not struct")
}

func TestFieldProperties(t *testing.T) {
	e := newElement(testValue, "")

	f := e.Fields()[1]
	assertEqual(t, "age", f.Name(), "Field name invalid")
	assertEqual(t, "20", f.Value(), "Field value invalid")
	assertEqual(t, "int", f.Type(), "Field type is not string")
	assertEqual(t, reflect.Int, f.Kind(), "Filed kind is not string")
}

func TestFieldsCount(t *testing.T) {
	e := newElement(testValue, "")
	fields := e.Fields()

	assertEqual(t, 2, len(fields), "Count of fields invalid")
}
