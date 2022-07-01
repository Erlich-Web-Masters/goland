package main

import "fmt"

func main() {

	var nombre string
	var edad int
	var altura float32

	fmt.Print("Ingresa tu nombre : ")
	fmt.Scanf("%s", &nombre)

	fmt.Print("Ingresa tu edad : ")
	fmt.Scanf("%d", &edad)

	fmt.Print("Ingresa tu altura : ")
	fmt.Scanf("%f", &altura)

	fmt.Printf("Tu nombre es %s tu edad es %d y tienes una altura de %f\n", nombre, edad, altura)

}
