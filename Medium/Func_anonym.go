package main

import "fmt"

func main() {
	ray := func() {
		fmt.Println("la funcion es anonima")
	}
	ray()
	func(x int) {
		x *= 4
		fmt.Println(x)
	}(9)

	fmt.Printf("%T\n", ray)

	Dani := funcio()
	i := Dani()
	fmt.Println(i)

	fmt.Println(funcio()())
}

func funcio() func() int {
	ray := func() int {
		fmt.Println("Hola Ray como estas")
		return 1
	}
	return ray
}
