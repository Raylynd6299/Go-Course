package main

import (
	"cmd/go/internal/par"
	"sync"
	"fmt"
)

func main() {
	par := make(chan int)
	impar := make(chan int)
	fanin := make(chan int)

	//enviar
	go enviarr(par, impar)

	//recibir
	go recibirr(par, impar, fanin)

	for v := range fanin {
		fmt.Println(v)
	}

	fmt.Println("Finalizando")
}

func enviarr(p, imp chan<- int) {
	for j := 0; j < 10; j++ {
		if j%2 == 0 {
			p <- j
		} else {
			imp <- j
		}
	} 
	close(p)
	close(imp)
}

func recibirr(p, imp, fanin <-chan int) {
	var wg sync.WaitGroup
	wg.Add(2)
	go fun(){
		for v:= range p{
			fanin <- v
		}
		wg.Done()
	}
	go fun(){
		for v:= range imp{
			fanin <- v
		}
		wg.Done()
	}
	wg.Wait()
	close(fanin)
}
