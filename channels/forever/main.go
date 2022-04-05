package main

import "fmt"

func main() {

	forever := make(chan int)

	go func() {
		for {

		}
	}()

	fmt.Println("Waiting forever!!!!!")
	<-forever
}
