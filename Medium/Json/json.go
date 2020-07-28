package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	//Marshall
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}

	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Read", "Ruby", "Maroon"},
	}

	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error", err)
	}
	//Para que la salida interprete los bytes usamos la Stdout, si no muestra el slice de bytes
	os.Stdout.Write(b)
	//UnMarshall
	var jsonBlob = []byte(`[
		{"Name":"Platypus","Order":"Monotremata"},
		{"Name":"Quoll","Order":"Dasyuromorphia"}
	]`)
	type Animal struct {
		Name  string
		Order string
	}

	var animal []Animal
	err = json.Unmarshal(jsonBlob, &animal)
	if err != nil {
		fmt.Println("error", err)
	}

	fmt.Printf("\n%+v\n", animal)

}
