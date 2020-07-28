package main

import "fmt"

var x int

func main() {
	fmt.Println(x)

	x++
	foo()
	fmt.Println(x)

	{
		y := 5
		fmt.Println(y)
	}
	//error abajo por las llaves, delimitan el scope de la variable y
	//fmt.Println(y)

	// cada que instancio genero una copia de todas las variables y las asociamos a la variable a y b, pero ambas son independientes
	a := incremento()
	fmt.Println(a(), a(), a(), a(), a())
	b := incremento()
	fmt.Println(b(), b(), b(), a())
}

func incremento() func() int {
	var l int
	return func() int {
		l++
		return l
	}
}

func foo() {
	fmt.Println(x)
	x++
}
