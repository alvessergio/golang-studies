package main

import (
	"fmt"
	"time"
)

func main() {

	hello := make(chan int)

	go func() {
		time.Sleep(time.Second * 2)
		hello <- 2
	}()

	select {
	case x := <-hello:
		fmt.Println(x)
	default:
		fmt.Println("default")
	}

	time.Sleep(time.Second * 5)
}
