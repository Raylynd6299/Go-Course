package main

import "fmt"

func main() {
	//unbuffer channel
	ca := make(chan int)

	go func() {
		ca <- 42
	}()
	fmt.Println(<-ca)

}
