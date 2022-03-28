package typeprinter

import (
	"io/ioutil"
	"os"
	"testing"
)

// Test Sprint.

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

type person struct {
	name string
	age  int
	work struct {
		post  string
		floor int
	}
}

func TestComplexSprint(t *testing.T) {
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
}`
	assertEqual(t, expected, real, "Sturct representation is invalid")
}

// Test non Sprint.

type kid struct {
	name string
	age  int
}

var kidValue = kid{name: "kid", age: 1}

var kidExpect = `kid {
	name: "kid"
	age: 1
}`

func TestSimpleSprintln(t *testing.T) {
	assertEqual(t, kidExpect+"\n", Sprintln(kidValue),
		"Sprintln simple struct representation is invalid")
}

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
