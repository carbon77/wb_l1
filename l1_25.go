package main

import (
	"fmt"
	"time"
)

func Sleep(duration time.Duration) {
	// Блокируем поток на время duration
	select {
	case <-time.After(duration):
	}
}

func main() {
	fmt.Println("Ждем 5 секунд")
	Sleep(5 * time.Second)
	fmt.Println("Прошло 5 секунд")
}
