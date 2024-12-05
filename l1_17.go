package main

import (
	"fmt"
	"sort"
)

// Функция бинарного поиска
// Принимает отсортированный массив и искомый элемент, возвращает
// индекс последнего искомого элемента или -1, если элемент не найден
func binarySearch(arr []int, target int) (int, bool) {
	low := 0
	high := len(arr)

	for low < high {
		// Находим середину массива
		mid := low + (high-low)/2

		// Сравниваем серединный элемент mid с искомым:
		// Если mid меньше искомого, то искомый элемент справа от среднего
		// Если mid больше искомого, то искомый элемент слева от среднего
		// Если mid равен искомому, то возвращаем индекс
		if arr[mid] < target {
			low = mid + 1
		} else if arr[mid] > target {
			high = mid
		} else {
			return mid, true
		}
	}

	return -1, false
}

func main() {
	arr := []int{5, 6, 1, 4, 20, 3, 0, 3}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	fmt.Printf("Массив: %v\n", arr)
	fmt.Printf("Введите искомый элемент:")
	var num int
	fmt.Scan(&num)
	idx, ok := binarySearch(arr, num)

	if !ok {
		fmt.Println("Элемент не найден")
	} else {
		fmt.Printf("Индекс элемента: %d\n", idx)
	}
}
