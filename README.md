# Struct printer
Package present struct in pretty and simple string.  

[API documentation](https://pkg.go.dev/github.com/backdround/structprinter)

### Installation
```bash
go get github.com/backdround/structprinter
```


### Example

```go
package main

import (
	"fmt"
	"github.com/backdround/structprinter"
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

	fmt.Print(structprinter.Sprint(p))
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
