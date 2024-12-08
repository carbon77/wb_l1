package main

import (
	"fmt"
	"log"
)

func main() {
	var a, b int64

	fmt.Print("Введите значение a (должно быть больше 2^20): ")
	fmt.Scan(&a)
	fmt.Print("Введите значение b (должно быть больше 2^20): ")
	fmt.Scan(&b)

	const threshold = 1 << 20
	if a <= threshold || b <= threshold {
		log.Fatalln("Числа должны быть больше 2^20!")
	}

	sum := a + b
	difference := a - b
	product := a * b

	// Проверка деления на 0
	var quotient float64
	if b != 0 {
		quotient = float64(a) / float64(b)
	} else {
		log.Fatalln("Деление на ноль!")
	}

	fmt.Printf("Сумма: %d\n", sum)
	fmt.Printf("Разность: %d\n", difference)
	fmt.Printf("Произведение: %d\n", product)
	fmt.Printf("Частное: %.2f\n", quotient)
}
