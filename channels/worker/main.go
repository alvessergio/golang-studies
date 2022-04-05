package main

import (
	"fmt"
	"time"
)

func worker(workerID int, msg chan int) {
	for response := range msg {
		fmt.Println("Worker: ", workerID, " Msg: ", response)
		time.Sleep(time.Second)
	}
}

func main() {

	msg := make(chan int)

	go worker(1, msg)
	go worker(2, msg)
	go worker(3, msg)

	for i := 0; i < 10; i++ {
		msg <- i
	}

}
