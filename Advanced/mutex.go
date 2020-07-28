package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	fmt.Println("Numero de CPU's", runtime.NumCPU())
	fmt.Println("Numero de GOroutines", runtime.NumGoroutine())
	contador := 0
	const gs = 100
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := contador
			v++
			runtime.Gosched() //Permite esperar para que las otra gorutines actuen
			//time.Sleep(time.Second * 2)
			contador = v
			mu.Unlock()
			wg.Done()
		}()
		fmt.Println("Numero de GOroutines", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println(contador)
}
