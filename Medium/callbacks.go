package main

import "fmt"

func main() {
	num := []int{1, 2, 3, 4, 5, 76, 87, 99, 0, 1, 4}
	sumPar(suma, num...)
	fmt.Println()
}

func sumPar(s func(xi ...int) int, xs ...int) {
	var y []int

	for _, v := range xs {
		if v%2 == 0 {
			y = append(y, v)
		}
	}

	fmt.Println("suma de pares", s(y...))
	fmt.Println("suma total", s(xs...))

}
func suma(xi ...int) int {
	res := 0
	for _, v := range xi {
		res += v
	}
	return res
}
