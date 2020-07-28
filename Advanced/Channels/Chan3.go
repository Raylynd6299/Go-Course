package main

import "fmt"

func main() {
	//buffer channel, puedes almacenar valores en mientras no llenes el buffer o no exista nadi que lo resiva
	ca := make(chan int, 2)

	ca <- 56
	ca <- 88
	fmt.Println(<-ca)
	fmt.Println(<-ca)
}
