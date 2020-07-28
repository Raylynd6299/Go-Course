package main

import "fmt"

type persona struct {
	nombre, apellido string
	edad             int
}

func (p *persona) cambiarNombre(nom string) {
	fmt.Println(p, p.nombre)
	p.nombre = nom
	fmt.Println(p, p.nombre)
}

func (p *persona) correr() {
	p.apellido = "Corredor"
	fmt.Println(p)
}

type maraton interface {
	correr()
}

func corre(m maraton) {
	m.correr()
}

func main() {
	a := 50
	b := &a
	x := []int{1, 2, 3, 4, 5, 6}
	y := &x
	fmt.Printf("%T  %T\n", a, y)
	fmt.Println(a, b)
	*y = nil
	fmt.Println(x, y)
	//Method Sets (uso de interfacez)
	//Los metodos que en su implementacion tienen como receptor a (t T) puedes pandarles como Parametro T y *T, sin tener um error,
	//Pero los que tiene un un receptor tipo (t *T) solo pueden recibir como parametro un *T

	//(t T) ---> T y *T
	//(t *T) ---> *T

	Ray := persona{
		nombre:   "Raymundo",
		apellido: "Pulido",
	}

	Ray.cambiarNombre("Ramon")
	fmt.Println(Ray)
	corre(&Ray)
	fmt.Println(Ray)

}
