package main

import "fmt"

func main() {
	//for {

	//}
	//for exp {

	//}
	//for int; ques; inc {

	//}
	x := []int{6, 5, 4, 8, 9, 5, 6, 1, 2}
	y := []int{9, 5, 5, 5, 3, 36}

	x = append(x, 9)
	x = append(x, y...)
	fmt.Println(cap(x))
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Println(x[0:2])
	fmt.Println(x)
	x = append(x[:2], x[4:]...)
	fmt.Println(x)
	z := make([]int, 10, 15)
	fmt.Println(z, cap(z), len(z))

	matriz := [][]int{z, x, y}
	fmt.Println(matriz)
}
