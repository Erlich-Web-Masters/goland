package main

import "fmt"

const (
	Domingo int = iota //Se le indica a go que las siguientes constantes son una secuencia de la constante Domingo que se inicializa con 0
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
)

func main() {
	fmt.Println(Domingo)
	fmt.Println(Lunes)
	fmt.Println(Martes)
	fmt.Println(Miercoles)
	fmt.Println(Jueves)
	fmt.Println(Viernes)
	fmt.Println(Sabado)
}
