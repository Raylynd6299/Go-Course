package main

import (
	"encoding/json"
	"fmt"
)

type persona struct {
	Nombre   string `json:"Nombre,omitempty`
	Apellido string `json:"Apellido"`
	edad     int    `json: "Edad"`
}

//los atributos con minusculas no son importados a Json, solo los que tienen Mayusculas al inicio se pasan
func main() {
	p1 := persona{
		Nombre:   "James",
		Apellido: "Bond",
		edad:     32,
	}
	p2 := persona{
		Nombre:   "Miss ",
		Apellido: "Money Penny",
		edad:     27,
	}
	personas := []persona{
		p1,
		p2,
	}
	fmt.Println(personas)

	sb, err := json.Marshal(personas)
	if err != nil {
		fmt.Println("Error ", err)
	}
	fmt.Println(string(sb))
}
