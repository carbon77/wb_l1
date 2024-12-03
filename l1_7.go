package main

import (
	"fmt"
	"sync"
	"time"
)

type ConcurrentMap[K comparable, V any] struct {
	data map[K]V
	mu   *sync.Mutex
}

func NewConcurrentMap[K comparable, V any]() *ConcurrentMap[K, V] {
	return &ConcurrentMap[K, V]{
		data: make(map[K]V),
		mu:   &sync.Mutex{},
	}
}

// Операция чтения потокобезопасна, поэтому не блокируем горутину
func (cm *ConcurrentMap[K, V]) Get(key K) (V, bool) {
	value, ok := cm.data[key]
	return value, ok
}

// Операция записи непотокобезопасна, поэтому необходаима блокировка, используем sync.Mutex
func (cm *ConcurrentMap[K, V]) Set(key K, value V) {
	cm.mu.Lock()
	cm.data[key] = value
	cm.mu.Unlock()
}

func main() {
	wg := &sync.WaitGroup{} // Для ожидания завершения горутин
	cm := NewConcurrentMap[int, int]()
	n := 10

	for i := 0; i < n; i++ {
		cm.Set(i, i)
	}

	// Запускаем две читающих из мапы горутины
	wg.Add(1)
	go func(cm *ConcurrentMap[int, int], wg *sync.WaitGroup) {
		for i := 0; i < n; i++ {
			num, _ := cm.Get(i)
			time.Sleep(1 * time.Second)
			fmt.Printf("Square: %d\n", num*num)
		}
		wg.Done()
	}(cm, wg)

	wg.Add(1)
	go func(cm *ConcurrentMap[int, int], wg *sync.WaitGroup) {
		for i := 0; i < n; i++ {
			num, _ := cm.Get(i)
			time.Sleep(1 * time.Second)
			fmt.Printf("Num + num: %d\n", num+num)
		}
		wg.Done()
	}(cm, wg)

	// Ожидаем завершение горутин
	wg.Wait()
}
