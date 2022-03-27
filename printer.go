package typeprinter

import (
	"fmt"
	"reflect"
)

func withPostfixOrAlternative(v string, postfix string, alternative string) string {
	if v != "" {
		return v + postfix
	}
	return alternative
}

func makeRepresentation(e element, indent string) string {
	switch e.Kind() {
	case reflect.Struct:
		return representStruct(e, indent)
	case reflect.String:
		return representString(e, indent)
	default:
		return representValue(e, indent)
	}
}

func representValue(e element, indent string) string {
	name := withPostfixOrAlternative(e.Name(), ": ", "")
	return fmt.Sprintf("%s%s%s", indent, name, e.Value())
}

func representString(e element, indent string) string {
	name := withPostfixOrAlternative(e.Name(), ": ", "")
	return fmt.Sprintf("%s%s\"%s\"", indent, name, e.Value())
}

func representStruct(e element, indent string) string {
	elementFields := e.Fields()

	// Get struct denotion.
	structName := ""
	if e.Name() != "" {
		structName = e.Name() + " "
	} else if e.Type() != "" {
		structName = e.Type() + " "
	}

	if len(elementFields) == 0 {
		return indent + structName + "{}"
	}

	representation := indent + structName + "{\n"
	for _, e := range elementFields {
		representation += makeRepresentation(e, indent+"\t") + "\n"
	}
	representation += indent + "}"

	return representation
}
