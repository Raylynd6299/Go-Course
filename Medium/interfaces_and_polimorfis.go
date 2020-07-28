package main

import "fmt"

type persona struct {
	nombre, apellido string
}

type agentesecreto struct {
	persona
	lpm bool
}
type humano interface {
	presentarse()
}

func main() {
	as1 := agentesecreto{
		persona: persona{nombre: "Raymudno", apellido: "Pulido"},
		lpm:     true,
	}
	defer fmt.Println(as1)
	as1.presentarse()
	p := persona{
		nombre:   "Ricardo",
		apellido: "perez",
	}
	bar(p)
	bar(as1)
}

func (a agentesecreto) presentarse() {
	fmt.Println("Hola, soy", a.nombre, " ", a.apellido, "el agente secreto se presento")
}

func (p persona) presentarse() {
	fmt.Println("Hola, soy", p.nombre, " ", p.apellido, "la persona se presento")
}

func bar(h humano) {
	//aserction
	switch h.(type) {
	case persona:
		fmt.Println("fui pasado a la funcion bar", h.(persona).nombre)
		break
	case agentesecreto:
		fmt.Println("fui pasado a la funcion bar", h.(agentesecreto).nombre)
		break

	}
	fmt.Println("fui pasado a la funcion bar", h)
}
