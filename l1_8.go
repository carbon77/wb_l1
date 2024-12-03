package main

import (
	"fmt"
	"math"
)

func main() {
	var num uint64
	fmt.Print("Введите число int64: ")
	fmt.Scan(&num)

	var i int
	fmt.Print("Введите номер бита для изменения (крайний правый бит имеет номер 1): ")
	fmt.Scan(&i)

	if i < 1 || i > 64 {
		panic("i должно быть на отрезке [1, 64]")
	}

	var value int
	fmt.Print("Введите число 1 или 0: ")
	fmt.Scan(&value)
	if value != 0 && value != 1 {
		panic("Неверное число!")
	}

	var newNum uint64
	if value == 1 {
		// Получаем маску вида 0...010...0 с помощью сдвига единицы влево на (i - 1)
		// Далее применяем логическое ИЛИ и меняем i-й бит на 1
		newNum = num | (1 << (i - 1))
	} else {
		// Получаем маску вида 1...101...1 с помощью формулы: 2 ** 64 - 1 - 2 ** (i - 1)
		// Далее применяем логическое И и меняем i-й бит на 0
		mask := uint64(math.Pow(2, 64)) - 1 - uint64(math.Pow(2, float64(i-1)))
		newNum = num & mask
	}

	fmt.Printf("New num: %d\n", newNum)
}
