package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	nums := []int{2, 4, 6, 8, 10}
	squares := make(chan int)
	wg := &sync.WaitGroup{}

	// Запускаем горутину для каждого числа
	for num := range nums {
		wg.Add(1)
		go findSquare(wg, squares, num)
	}

	// Ждем завершения всех горутин и закрываем канал
	go func() {
		wg.Wait()
		close(squares)
	}()

	// Читаем данные из канала и выводим результат
	for square := range squares {
		time.Sleep(1 * time.Second)
		fmt.Println(square)
	}
}

func findSquare(wg *sync.WaitGroup, squares chan int, num int) {
	defer wg.Done()
	squares <- num * num
}
