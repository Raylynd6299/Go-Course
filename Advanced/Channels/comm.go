package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		c <- 42
		close(c)
	}()
	v, ok := <-c
	fmt.Println("valores", v, ok)
	v, ok = <-c
	fmt.Println("valores", v, ok)

}

func enviarr(p, imp, sal chan<- int) {
	for j := 0; j < 100; j++ {
		if j%2 == 0 {
			p <- j
		} else {
			imp <- j
		}
	}
	close(sal)
}

func recibirr(p, imp, sal <-chan int) {
	for {
		select {
		case v := <-p:
			fmt.Println("Desde el canal Par", v)
		case v := <-imp:
			fmt.Println("Desde el canal imPar", v)
		case i, ok := <-sal:
			if !ok {
				fmt.Println("desde comma OK", i, ok)
			} else {
				fmt.Println("desde ok false")
			}
			return
		}
	}
}
