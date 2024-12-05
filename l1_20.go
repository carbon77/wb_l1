package main

import (
	"fmt"
	"strings"
)

func ReverseWords(s string) string {
	words := strings.Split(s, " ")
	reversed := make([]string, len(words))

	for i, word := range words {
		reversed[len(words)-i-1] = word
	}

	return strings.Join(reversed, " ")
}

func main() {
	sentence := "snow dog sun    hello"

	fmt.Printf("Initial string: %s\n", sentence)
	fmt.Printf("String with reversed words: %s\n", ReverseWords(sentence))
}
