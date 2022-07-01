package main

import "fmt"

func main() {

	//Explicitamente
	var nombre string = "Airam"
	var edad int = 8

	//Implicitamente
	color := "Azul"
	var longitud = 1.85

	//multiples variables
	var base, ancho, altura float64
	user, pass, email := "Jose", "Jose123", "jose@gamil.com"

	fmt.Println("Nombre :", nombre)
	fmt.Println("Edad :", edad)
	fmt.Println("Color :", color)
	fmt.Println("Longitud :", longitud)
	fmt.Println(base, ancho, altura)
	fmt.Println(user, pass, email)

	//Nota: Todas las varibles vacias se inician con su valor 0 o string vacios

}
