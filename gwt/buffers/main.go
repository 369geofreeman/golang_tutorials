package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println("Using 'bytes' standard package")

	var b bytes.Buffer
	(&b).Write([]byte("Hello World\n"))
	fmt.Fprintf(&b, "Holy %s\n", "smokes batman!")
	fmt.Println(b.String())
}
