package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	w := &sync.WaitGroup{}
	// Способ 1: с помощью сигнального канала
	s := make(chan struct{})

	w.Add(1)
	go func() {
		fmt.Println("Goroutine with signal channel")
		for {
			select {
			case <-s:
				fmt.Println("Stoping with signal channel")
				w.Done()
				return
			}
		}
	}()

	// Способ 2: с помощью Context
	ctx, cancel := context.WithCancel(context.Background())
	w.Add(1)
	go func() {
		fmt.Println("Goroutine with context")
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Stoping with context")
				w.Done()
				return
			}
		}
	}()

	// Способ 3: с помощью time.After
	w.Add(1)
	go func() {
		fmt.Println("Goroutine with time.After")
		for {
			select {
			case <-time.After(3 * time.Second):
				fmt.Println("Stoping with time.After")
				w.Done()
				return
			}
		}
	}()

	// Способ 4: с помощью флага
	w.Add(1)
	stop := false
	go func() {
		fmt.Println("Goroutine with flag")
		for !stop {
		}
		fmt.Println("Stoping with flag")
		w.Done()
	}()

	time.Sleep(5 * time.Second)
	s <- struct{}{}
	cancel()
	stop = true
	w.Wait()
}
