package main

import "fmt"

func main() {

	// numeros := []int{1, 2, 3, 4}

	// numeros = append(numeros, 5)
	// numeros = append(numeros, 6)
	// numeros = append(numeros, 7)

	// otrosNumeros := numeros[0:3]

	// fmt.Println(numeros)
	// fmt.Println(otrosNumeros)

	meses := []string{"Enero", "Febrero", "Marzo", "Abril", "Mayo"}

	//Un puntero
	//Una longitud
	//una capacidad

	longitud := len(meses)
	capacidad := cap(meses)

	meses = append(meses, "Junio") // Si la estructura se encuentra en su capacidad máxima generará otro arreglo

	longitud = len(meses)
	capacidad = cap(meses)

	fmt.Printf("La longitus es : %v, la capacidad es : %v %p\n", longitud, capacidad, meses)
	fmt.Println(meses)
	//%p dice la direccion en memoria de meses

	//Nota los slice son una referencia o acceso a una subsecuencia de elementos de un arreglo
}
