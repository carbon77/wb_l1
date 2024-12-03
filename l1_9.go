package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	nums := []int{5, 1, 6, 10, 3, 5}

	// Горутина для записи чисел в поток ch1
	go func() {
		for _, num := range nums {
			time.Sleep(1 * time.Second)
			ch1 <- num
		}
		close(ch1)
	}()

	// Горутина для чтения чисел из потока ch1
	go func() {
		for num := range ch1 {
			ch2 <- num * 2
		}
	}()

	// Читаем числа из потока ch2
	for num := range ch2 {
		fmt.Println(num)
	}
}
