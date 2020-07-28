package main

import (
	"encoding/json"
	"fmt"
)

type persona struct {
	Nombre   string
	Apellido string
	edad     int
}

//los atributos con minusculas no son importados a Json, solo los que tienen Mayusculas al inicio se pasan
func main() {
	data := []byte(`[{"Nombre":"James","Apellido":"Bond"},
	{"Nombre":"Miss ","Apellido":"Money Penny"}]`)
	var personas []persona
	err := json.Unmarshal(data, &personas)
	if err != nil {
		fmt.Println("Error ", err)
	}
	fmt.Println(personas)
	fmt.Printf("%T", personas)

	types := []interface{}{"a", 6, 6.0, true, personas} // Slice Generico
	for _, v := range types {
		fmt.Printf("%T\n", v)
	}

	fmt.Println(json.Valid(data)) //Checa si el Json es valido

}
