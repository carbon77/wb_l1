package main

import (
	"fmt"
	"sync"
	"time"
)

// Структура счетчика
type Counter struct {
	value int
	mu    *sync.Mutex
}

func NewCounter() *Counter {
	return &Counter{
		value: 0,
		mu:    &sync.Mutex{},
	}
}

// Операция инкремента не потокобезопасна, поэтому необходима блокировка. Используем sync.Mutex
func (c *Counter) Increment() {
	defer c.mu.Unlock()
	c.mu.Lock()
	c.value++
}

// Операция чтения потокобезопасна, поэтому блокировка не нужна
func (c *Counter) Get() int {
	return c.value
}

func main() {
	wg := &sync.WaitGroup{}
	counter := NewCounter()

	// Запускаем 10 горутин, каждая из которых увеличивает счетчик на 50
	for j := 0; j < 10; j++ {
		wg.Add(1)
		go func() {
			fmt.Printf("Goroutine %d started\n", j+1)
			defer wg.Done()
			time.Sleep(5 * time.Second)
			for i := 0; i < 50; i++ {
				counter.Increment()
			}
		}()
	}

	// Ожидаем завершения всех горутин
	wg.Wait()
	fmt.Printf("Counter value = %d\n", counter.Get())
}
