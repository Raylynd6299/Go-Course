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

# Variables \n #
```Go
#Explicita \n
  var Nombre Tipo
#Implicito\n
  Nombre := Valor
```
# Slice en Go #
```Go
#Creacion de Slice
#Form 1
  array := []int{var1,var2,var3}
#Form 2
  array := make([]int,longitud,capacidad)

#Versiones de For
  for{

  }
  end
#Version comun del For
  for i:=0; i<max ; i++ {

  }
  end
#Version de For Each
  for in,v := range Slice{

  }
  end


#Agregar elementos al slice
  array = append(array,elemento)

#Recorrer un slice por partes
  array[1:5]