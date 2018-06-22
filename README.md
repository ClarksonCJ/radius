# radius [![GoDoc](https://godoc.org/github.com/ClarksonCJ/radius?status.svg)](https://godoc.org/github.com/ClarksonCJ/radius)

a Go (golang) [RADIUS](https://tools.ietf.org/html/rfc2865) client and server implementation

## Installation

    go get -u github.com/ClarksonCJ/radius

## Client example

```go
package main

import (
	"context"
	"fmt"

	"github.com/ClarksonCJ/radius"
	. "github.com/ClarksonCJ/radius/rfc2865"
)

func main() {
	packet := radius.New(radius.CodeAccessRequest, []byte(`secret`))
	UserName_SetString(packet, "tim")
	UserPassword_SetString(packet, "12345")
	response, err := radius.Exchange(context.Background(), packet, "localhost:1812")
	if err != nil {
		panic(err)
	}

	if response.Code == radius.CodeAccessAccept {
		fmt.Println("Accepted")
	} else {
		fmt.Println("Denied")
	}
}
```

## RADIUS Dictionaries

Included in this package is the command line program `radius-dict-gen`. It can be installed with:

    go get -u github.com/ClarksonCJ/radius/cmd/radius-dict-gen

Given a FreeRADIUS dictionary, the program will generate helper functions and types for reading and manipulating RADIUS attributes in a packet. It is recommended that generated code be used for any RADIUS dictionary you would like to consume.

Included in this repository are sub-packages of generated helpers for commonly used RADIUS attributes, including [`rfc2865`](https://godoc.org/github.com/ClarksonCJ/radius/rfc2865) and [`rfc2866`](https://godoc.org/github.com/ClarksonCJ/radius/rfc2866).

## License

MPL 2.0
