package main

import (
	"fmt"
	"sort"
)

type Persona struct {
	Nombre string
	Edad   int
}

func (p Persona) String() string {
	return fmt.Sprintf("%s:%d", p.Nombre, p.Edad)
}

type Byage []Persona

func (a Byage) Len() int           { return len(a) }
func (a Byage) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Byage) Less(i, j int) bool { return a[i].Edad < a[j].Edad }

func main() {
	ss := []int{3, 5, 2, 5, 8, 9, 0, 11, 9}
	cads := []string{"Hooa", "Hola", "Soy ray", "Ricardo", "Daniela"}
	fmt.Println(ss)
	sort.Ints(ss)
	fmt.Println(ss)
	fmt.Println(cads)
	sort.Strings(cads)
	fmt.Println(cads)

	p1 := Persona{"Edgar", 32}
	p2 := Persona{"Mariana", 18}
	p3 := Persona{"Carolina", 25}
	p4 := Persona{"Adriana", 31}
	p5 := Persona{"Dulce", 30}
	p6 := Persona{"Raymundo", 21}

	personas := []Persona{p1, p2, p3, p4, p5, p6}

	fmt.Println(personas)
	sort.Sort(Byage(personas))
	fmt.Println(personas)

}
