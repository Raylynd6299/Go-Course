package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Numero de CPU's", runtime.NumCPU())
	fmt.Println("Numero de GOroutines", runtime.NumGoroutine())
	var contador int64 = 0
	const gs = 100
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&contador, 1)
			runtime.Gosched() //Permite esperar para que las otra gorutines actuen
			//time.Sleep(time.Second * 2)
			fmt.Println("Contador", atomic.LoadInt64(&contador))
			wg.Done()
		}()
		fmt.Println("Numero de GOroutines", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println(contador)
}
