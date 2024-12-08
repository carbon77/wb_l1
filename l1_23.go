package main

import (
	"fmt"
	"log"
)

func RemoveFromSlice(slice []int, idx int) []int {
	return append(slice[:idx], slice[idx+1:]...)
}

func main() {
	slice := []int{3, 5, 1, 3, 7, 2, 4, 1}
	var idx int

	fmt.Printf("Начальный слайс: %v\n", slice)
	fmt.Print("Введите индекс элемента для удаления: ")
	fmt.Scan(&idx)

	if 0 > idx || idx >= len(slice) {
		log.Fatalf("Допустимые значения: [%d, %d]\n", 0, len(slice)-1)
	}

	slice = RemoveFromSlice(slice, idx)
	fmt.Printf("Слайс без %d элемента: %v\n", idx, slice)
}
