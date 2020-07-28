package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPU's\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	wg.Add(1) //Numero de Goroutines
	go foo()
	bar()

	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPU's\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Wait() // Espera a que todas las goroutines terminen
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo", i)
	}
	wg.Done() // Le dice que ya acbo la Goroutines
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar", i)
	}
}
