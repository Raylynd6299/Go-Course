package main

import (
	"fmt"
)

type complejo int

//Creacion explicita de Variables
// var identtificador Tipo_de_dato 
var z complejo
var y int = 9
var fll = 9.0

func main() {
	fmt.Println("Hola Mundo")
	fmt.Println("Control de Salida de Datos")
	fmt.Println(z, y, fll)
	
	//Mostrar el tipo de dato
	fmt.Printf("%T", z)
	z = 9
	y = int(z)
}
