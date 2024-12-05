package main

import (
	"fmt"
	"strings"
)

var justString string

func createString(length int) string {
	return strings.Repeat("A", length)
}

func someFunc() {
	justString = createString(100)
}

func main() {
	someFunc()
	fmt.Println(justString)
}
