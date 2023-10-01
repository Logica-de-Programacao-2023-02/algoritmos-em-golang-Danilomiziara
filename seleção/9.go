package main

import (
	"fmt"
)

// Faça um algoritmo que leia três números reais e mostre-os em ordem crescente.
func main() {
	var x, y, z float64
	fmt.Println("digite 3 numeros")
	fmt.Scan(&x, &y, &z)
	if x >= y && y >= z {
		fmt.Println(x, y, z)

	} else if y >= x && x >= z {
		fmt.Println(y, x, z)

	} else if z >= x && x >= y {
		fmt.Println(z, x, y)
	} else if z >= y && y >= x {
		fmt.Println(z, y, x)
	} else if x >= z && z >= y {
		fmt.Println(x, z, y)
	} else if y >= z && z >= x {
		fmt.Println(y, z, x)
	} else {

	}

}

// COLOQUEI EM ORDEM DECRESCENTE SEM QUERER, SE VIRA
