package structprinter

import (
	"testing"
	"reflect"
)

func assertEqual(t *testing.T, real interface{}, expected interface{},
	failMessage string) {
	if expected != real {
		t.Errorf("%s\nexpected:%v\nreal:%v\n", failMessage, expected, real)
	}
}

type testStruct struct {
	name string
	age  int
}
var testValue = testStruct{"bob", 20}

func TestRootProperties(t *testing.T) {
	e := newElement(testValue, "bob-element")

	assertEqual(t, e.Name(), "bob-element", "Root name invalid")
	assertEqual(t, e.Value(), "{bob 20}", "Root value invalid")
	assertEqual(t, e.Type(), "testStruct", "Root type invald")
	assertEqual(t, e.Kind(), reflect.Struct, "Struct kind is not struct")
}

func TestFieldProperties(t *testing.T) {
	e := newElement(testValue, "")

	f := e.Fields()[1]
	assertEqual(t, f.Name(), "age", "Field name invalid")
	assertEqual(t, f.Value(), "20", "Field value invalid")
	assertEqual(t, f.Type(), "int", "Field type is not string")
	assertEqual(t, f.Kind(), reflect.Int, "Filed kind is not string")
}

func TestFieldsCount(t *testing.T) {
	e := newElement(testValue, "")
	fields := e.Fields()

	assertEqual(t, len(fields), 2, "Count of fields invalid")
}
