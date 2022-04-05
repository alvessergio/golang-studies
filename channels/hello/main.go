package main

import "fmt"

//T01
func main() {

	hello := make(chan string)

	//T02
	go func() {
		hello <- "Hello"
	}()

	result := <-hello

	fmt.Println("Result: ", result)
}
