package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Worker struct {
	Index int       // Идентификатор воркера
	Quit  chan bool // Поток для завершения воркера
	Data  chan int  // Поток с данными
	Wg    *sync.WaitGroup
}

func (w *Worker) Run() {
	for {
		select {
		case <-w.Quit: // Завершение ворверка
			fmt.Printf("Worked %d closing\n", w.Index)
			w.Wg.Done()
			return
		case num := <-w.Data:
			fmt.Printf("Worker %d = %d\n", w.Index, num)
		}
	}
}

func main() {
	// Поток сигналов для "чистого" завершения программы
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	var nWorkers int

	fmt.Print("Enter number of workers: ")
	fmt.Scan(&nWorkers)

	dataChan := make(chan int)
	wg := &sync.WaitGroup{}
	var workers []*Worker

	// Запуск воркеров
	for i := 0; i < nWorkers; i++ {
		quit := make(chan bool)
		worker := Worker{i, quit, dataChan, wg}
		wg.Add(1)
		go worker.Run()
		workers = append(workers, &worker)
	}

	for {
		time.Sleep(1 * time.Second)
		select {
		case <-sigChan:
			// Выход из программы, завершение горутин и закрытие потока данных
			fmt.Println("Closing...")
			for _, worker := range workers {
				worker.Quit <- true
			}
			wg.Wait()
			fmt.Println("Closing data channel")
			close(dataChan)
			return
		default:
			// Отправка данных в поток данных
			dataChan <- rand.Int()
		}
	}
}
