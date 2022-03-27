// Package typeprinter present struct in pretty and simple string.
package typeprinter


// Sprint return string with a simple struct representation.
func Sprint(v interface{}) string {
	if v == nil {
		return "\n"
	}
	return makeRepresentation(newElement(v, ""), "")
}

