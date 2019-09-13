package main

import (
	"fmt"
)

func foo(s string) string {
	return "test-gorel:" + s
}

func main() {
	fmt.Println(foo("check"))
}
