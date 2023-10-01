package main

import "fmt"

// Faça um algoritmo que leia um número inteiro positivo e mostre todos os seus divisores.
func main() {
	var x int
	fmt.Println("diga um número inteiro")
	fmt.Scan(&x)
	for i := 1; i <= x; i++ {
		if x%i == 0 {
			fmt.Printf("divisor de %d: %d\n", x, i)

		}

	}

}
