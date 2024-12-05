package main

import (
	"fmt"
	"strings"
)

func CheckUnique(s string) bool {
	// Создаем структуру данных множество с помощью map, где ключ это элемент множества
	set := make(map[rune]bool)
	runes := []rune(strings.ToLower(s))

	for _, rune := range runes {
		// Если символ уже есть в множестве, то условию уникальность нарушено, выходим из функции
		if set[rune] {
			return false
		}
		set[rune] = true
	}

	// Если прошли цикл, то все символы в строке уникальны
	return true
}

func main() {
	tests := []string{"abcd", "abCdefAaf", "aabcd"}

	for i, t := range tests {
		fmt.Printf("Test %d:\n", i+1)
		fmt.Printf("\t%s --> %v\n", t, CheckUnique(t))
	}
}
