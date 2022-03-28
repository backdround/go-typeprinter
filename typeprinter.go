// Package typeprinter present struct in pretty and simple string.
package typeprinter

import "fmt"

// Sprint make simple representation from passed value.
func Sprint(v interface{}) string {
	if v == nil {
		return ""
	}
	return makeRepresentation(newElement(v, ""), "")
}

// Sprintln make simple representation from passed value.
// New line is appended.
func Sprintln(v interface{}) string {
	return Sprint(v) + "\n"
}

// Print is printing simple value representation to stdout.
func Print(v interface{}) {
	fmt.Print(Sprint(v))
}

// Println is printing simple value representation to stdout.
// New line is appended.
func Println(v interface{}) {
	fmt.Println(Sprint(v))
}
