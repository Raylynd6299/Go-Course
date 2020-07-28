package main

import (
	"fmt"
)

const v1 int = 69

const (
	ray  = "Raymundo"
	dani = "Daniela"
)

func main() {
	s1 := "Hola Mundo"
	s2 := `Esta linea no esta partida
	 o si`

	fmt.Println(s1)
	fmt.Printf("Es de tipo %T\n", s1)

	fmt.Println(s2)
	fmt.Printf("Es de tipo %T\n", s2)

	bs := []byte(s1)
	fmt.Println(bs)
	fmt.Printf("Es de tipo %T\n", bs)
	fmt.Println(v1)
	a := `here is something
	as 
	a 
	raw string
	literal
	"you see"
	another thing`

	fmt.Println(a)
}
