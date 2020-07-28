package main

import (
	"fmt"
)

type persona struct {
	nombre   string
	apellido string
	edad     int
}

type agentesecreto struct {
	persona
	lpm bool
}

func main() {
	Ray := persona{nombre: "Raymundo", apellido: "Pulido"}

	Dani := persona{"Angela", "Banderas", 21}

	jb := agentesecreto{
		persona{
			"james",
			"Bond",
			32,
		},
		true,
	}

	Jobs := struct {
		nombre   string
		apellido string
		altura   string
	}{
		"Joanna",
		"Aguilar",
		"160",
	}

	fmt.Println(Jobs)

	fmt.Println(jb)
	fmt.Println(jb.nombre)
	fmt.Println(Ray, Dani)
}
