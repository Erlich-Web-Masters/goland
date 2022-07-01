package main

import (
	"fmt"
	"reflect"
)

func main() {
	frase := "Nunca pares de aprender" // Cadana de caracteres

	longitud := len(frase) //devuelve un int / Longitud de un string

	fmt.Println(longitud)
	fmt.Println(reflect.TypeOf(longitud)) //Conocer el tipo de dato

	primerCaracter := frase[0]
	ultimoCaracter := frase[len(frase)-1]

	fmt.Printf("%c\n", primerCaracter) // %c para darle formato a la varibale que por defecto devuelve codigo aski
	fmt.Printf("%c\n", ultimoCaracter)

}
