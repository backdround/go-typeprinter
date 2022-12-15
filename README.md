[![Go Reference](https://img.shields.io/badge/go-reference-%2300ADD8?style=flat-square)](https://pkg.go.dev/github.com/backdround/go-typeprinter)
[![Tests](https://img.shields.io/github/actions/workflow/status/backdround/go-typeprinter/tests.yml?branch=main&label=tests&style=flat-square)](https://github.com/backdround/go-typeprinter/actions)
[![Codecov](https://img.shields.io/codecov/c/github/backdround/go-typeprinter?style=flat-square)](https://app.codecov.io/gh/backdround/go-typeprinter/)
[![Go Report](https://goreportcard.com/badge/github.com/backdround/go-typeprinter?style=flat-square)](https://goreportcard.com/report/github.com/backdround/go-typeprinter)

# Type printer
Package present type in pretty and simple string.  

### Installation
```bash
go get github.com/backdround/go-typeprinter
```


### Example

```go
package main

import (
	"fmt"
	"github.com/backdround/go-typeprinter"
)

type Person struct {
	name string
	age  int
	work struct {
		post  string
		floor int
	}
}

func main() {
	p := Person{
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

	fmt.Print(typeprinter.Sprint(p))
}
```

Output:
```
Person {
	name: "bob"
	age: 20
	work {
		post: "boss"
		floor: 32
	}
}
```
