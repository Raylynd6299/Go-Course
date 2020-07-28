package main

import "fmt"

type persona struct {
	nombre, apellido string
}

type agentesecreto struct {
	persona
	lpm bool
}

func main() {
	as1 := agentesecreto{
		persona: persona{nombre: "Raymudno", apellido: "Pulido"},
		lpm:     true,
	}
	defer fmt.Println(as1)
	as1.presentarse()
}

func (a agentesecreto) presentarse() {
	fmt.Println("Hola, soy", a.nombre, " ", a.apellido)
}
