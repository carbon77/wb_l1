package main

import "fmt"

// Так как строка может иметь символы unicode и большинство имеет развер более 1 байта, необходимо работать с типом rune
func ReverseString(s string) string {
	runes := []rune(s)
	reversed := make([]rune, len(runes))

	for i, rune := range runes {
		reversed[len(runes)-i-1] = rune
	}
	return string(reversed)
}

func main() {
	s := "главрыба こんにちは"
	fmt.Printf("Initial string: %s\n", s)
	fmt.Printf("Reversed string: %s\n", ReverseString(s))
}
