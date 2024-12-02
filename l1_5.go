package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var nSeconds int

	fmt.Print("Enter n of seconds: ")
	fmt.Scan(&nSeconds)

	var nums = make(chan int)
	var quitReader = make(chan bool)
	var quitWriter = make(chan bool)

	go func() {
		for {
			select {
			case num := <-nums:
				fmt.Printf("Read num: %d\n", num)
			case <-quitReader:
				fmt.Println("Quit reader")
				return
			}
		}
	}()

	go func() {
		for {
			select {
			case <-quitWriter:
				fmt.Println("Quit writer")
				return
			default:
				time.Sleep(1 * time.Second)
				var num = rand.Int()
				fmt.Printf("Sent num: %d\n", num)
				nums <- num
			}
		}
	}()

	time.Sleep(time.Duration(nSeconds) * time.Second)
	quitWriter <- false
	quitReader <- false
	close(nums)
	close(quitWriter)
	close(quitReader)
}
