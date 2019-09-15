// Copyright (c) 2019 hankei6km
// Licensed under the MIT License. See LICENSE in the project root.
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
