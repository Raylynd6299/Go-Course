# Go-Course

## Repositorio del Curso de Go, con pequeños programas ejemplo de cada tema ##

Go es un lenguaje de programación concurrente y compilado inspirado en la sintaxis de C, que intenta ser dinámico como Python y con el rendimiento de C o C++. Ha sido desarrollado por Google, y sus diseñadores iniciales fueron Robert Griesemer, Rob Pike y Ken Thompson.

# Hello world #
```Go
#Form 1
  fmt.Println("Hola Mundo")
#Form 2
  fmt.Printf("Hola Mundo")
```

# Variables #
```Go
#Explicita \n
  var Nombre Tipo
#Implicito\n
  Nombre := Valor
```
#Tipos de Datos#
```Go
#Enteros
  var entero int 
#Flotante
  var real float64
#Cadenas
  var cadena string
#Slice o Arreglo
  var array []int //int se cambia por el tipo deseado
#Matrices
  var matris [][]int
#Maps
  var mapp map[Tipo_llave]Tipo_Valor
#Estructuras
  type NombreEstructura struct{
  	Valores
  }
#Interfaces
  type NombreInterfaz interface {
  	NombreFuncion Retorno
  	...
  }
```
# Slice en Go #
```Go
#Creacion de Slice
#Form 1
  array := []int{var1,var2,var3}
#Form 2
  array := make([]int,longitud,capacidad)
```
# Ciclos de Control #
```
# Versiones de For 
  for{

  }
#Version comun del For
  for i:=0; i<max ; i++ {

  }
#Version de For Each
  for in,v := range Slice{

  }

#Agregar elementos al slice
  array = append(array,elemento)

#Recorrer un slice por partes
  array[1:5]
```
# Estructura de Funciones #

```
# La estructura 
   func (r receptor) NameFucntion(Argumentos) (Valores de Retorno){}
   
#Ejemplo
   func Suma(x,y int) int{
   	return x+y
   }
#Ejemplo 2
  type Humano struct{
  	Nombre,Apellido string
  	Edad int
  }
  func (h Humano) Presentar() string{
  	return fmt.Sprint("Hola, yo soy",h.Nombre," ",h.Apellido)
  }
```
