package main

import "fmt"

func main() {

	numeros := []int{1, 2, 3, 4}

	numeros = append(numeros, 5)
	numeros = append(numeros, 6)
	numeros = append(numeros, 7)

	otrosNumeros := numeros[0:3]

	fmt.Println(numeros)
	fmt.Println(otrosNumeros)
}
