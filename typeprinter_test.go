package typeprinter

import (
	"io/ioutil"
	"os"
	"testing"
)

// Simple api test.

type kid struct {
	name string
	age  int
}

var kidValue = kid{name: "kid", age: 1}

var kidExpect = `kid {
	name: "kid"
	age: 1
}`

// interceptStdout execute f, intercept stdout output and return it like string.
func interceptStdout(f func()) string {
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}

	oldStdout := os.Stdout
	os.Stdout = w

	go func() {
		f()
		os.Stdout.Close()
	}()

	data, err := ioutil.ReadAll(r)
	os.Stdout = oldStdout

	if err != nil {
		panic(err)
	}

	return string(data)
}

func TestSimpleSprint(t *testing.T) {
	assertEqual(t, kidExpect, Sprint(kidValue),
		"Sprint simple struct representation is invalid")
}

func TestSimpleSprintln(t *testing.T) {
	assertEqual(t, kidExpect+"\n", Sprintln(kidValue),
		"Sprintln simple struct representation is invalid")
}

func TestSimplePrint(t *testing.T) {
	output := interceptStdout(func() {
		Print(kidValue)
	})

	assertEqual(t, kidExpect, output,
		"Print simple struct representation output is invalid")
}

func TestSimplePrintln(t *testing.T) {
	output := interceptStdout(func() {
		Println(kidValue)
	})

	assertEqual(t, kidExpect+"\n", output,
		"Println simple struct representation output is invalid")
}

// Complex struct.
type person struct {
	name string
	age  int
	work struct {
		post  string
		floor int
	}
}

func TestComplexSprint(t *testing.T) {
	p := person{}
	p.name = "bob"
	p.age = 20
	p.work.post = "boss"
	p.work.floor = 32

	real := Sprint(p)

	expected :=
		`person {
	name: "bob"
	age: 20
	work {
		post: "boss"
		floor: 32
	}
}`
	assertEqual(t, expected, real, "Sturct representation is invalid")
}
