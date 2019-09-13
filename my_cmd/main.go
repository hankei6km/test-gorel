package main

import (
	"fmt"
)

var version = ""

func foo(s string) string {
	return "test-gorel: " + s
}

func main() {
	fmt.Println(foo("check"))
	fmt.Println("version:", version)
}
