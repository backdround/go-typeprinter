package typeprinter

import "testing"

func assertEqual(t *testing.T, expected interface{}, real interface{},
	failMessage string) {
	t.Helper()
	if expected != real {
		t.Errorf("%s\nexpected:%v\nreal:%v\n", failMessage, expected, real)
	}
}
