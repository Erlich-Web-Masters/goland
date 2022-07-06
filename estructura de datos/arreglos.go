package main

import "fmt"

func main() {

	var numeros [5]int // number el 5 detona la cantidad de elementos

	numeros[0] = 100

	number := [5]int{100, 200, 300, 400, 500}
	//number := [...]int{100, 200, 300, 400, 500} cuando se quiere que el compilador defina la longitud del array

	occidente := [...]string{"Habana", "Pinar del Rio", "Matazas", "Artemisa", "Mayabeque"}

	//Sub Arrglo o slice

	otroOccidente := occidente[0:2]

	fmt.Println(numeros)
	fmt.Println(number)
	fmt.Println(occidente)
	fmt.Println(otroOccidente)

}
