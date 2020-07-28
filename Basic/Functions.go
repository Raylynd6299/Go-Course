package main

import (
	"fmt"
)

//func (r receptor) identificador(parametros) retorno(s) {codigo}

func main() {
	foo()
	
	//Defer hace que lo que esta a la derecha se ejecutara al final de la Funcion
	defer bar("Ray")
	
	s1 := woo("Dani")
	fmt.Println(s1)
	x, y := saludar("Hola", "mundo")

	fmt.Println(x, y)
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	foo2("Hola", 3, 4, 5, 6, 8)
	foo2(x, slice...)
}

func foo() {
	fmt.Println("Hola desde foo")
}

func bar(s string) {
	fmt.Println("Hola", s)
}

func woo(s string) string {
	return fmt.Sprint("Hola desde woo, ", s)
}

func saludar(n, a string) (string, bool) {
	p := fmt.Sprint(n, " ", a, ` dice "hola".`)
	q := true
	return p, q
}
func foo2(s string, x ...int) {
	fmt.Printf("%T\n", x)
	fmt.Println(s)
	res := 0
	for _, v := range x {
		res += v
	}
	fmt.Println(res)

}
