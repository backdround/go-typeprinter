package typeprinter

import (
	"fmt"
	"reflect"
)

// element is used for simple traversion over structure fields.
type element struct {
	name string

	rType  reflect.Type
	rValue reflect.Value
}

func newElement(v interface{}, name string) element {
	return element{
		name:   name,
		rType:  reflect.TypeOf(v),
		rValue: reflect.ValueOf(v),
	}
}

func (e element) Fields() []element {
	fieldsNum := e.rValue.NumField()
	fields := make([]element, 0, fieldsNum)

	for i := 0; i < fieldsNum; i++ {

		fieldStruct := e.rType.Field(i)

		fieldElement := element{
			name:   fieldStruct.Name,
			rType:  fieldStruct.Type,
			rValue: e.rValue.Field(i),
		}

		fields = append(fields, fieldElement)
	}
	return fields
}

func (e element) Name() string {
	return e.name
}

func (e element) Value() string {
	return fmt.Sprintf("%v", e.rValue)
}

func (e element) Type() string {
	return e.rType.Name()
}

func (e element) Kind() reflect.Kind {
	return e.rType.Kind()
}
