package main

import "fmt"

func main() {
	a := 5
	b := 3

	fmt.Printf("a=%d, b=%d\n", a, b)
	a, b = b, a
	fmt.Printf("a=%d, b=%d\n", a, b)
}
