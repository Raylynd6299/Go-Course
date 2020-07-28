package main

import (
	"fmt"
)

func main() {
	//llave-valor
	//map[Tipo_de_llave]Tipo_valor{ pares }
	m := map[string]int{
		"Batman": 32,
		"Robin":  27,
	}
	fmt.Println(m)
	fmt.Println(m["Robin"])

	v, ok := m["Ray"]
	fmt.Println(v, ok)

	if v, ok := m["Batman"]; ok {
		fmt.Println("Existe el valor", v)
	}
	// ok, regresa si existe o no en booleano

	//agregamos el valor
	m["Ray"] = 21
	//for each
	for llave, val := range m {
		fmt.Println(llave, val)
	}

	//Borrar par valor
	delete(m, "Robin")
	for llave, val := range m {
		fmt.Println(llave, val)
	}

	if v, ok := m["Batman"]; ok {
		fmt.Println("Se borro la llave", v)
		delete(m, "Batman")
	}
}
