package main

import "fmt"

// Faça um algoritmo que leia dois números inteiros e mostre o maior deles.
func main() {
	var x int
	var y int
	fmt.Println("diga dois números inteiros")
	fmt.Scan(&x, &y)
	if x > y {
		fmt.Println(x, "é maior que", y)
	} else if x == y {
		fmt.Println("os 2 números são iguais")
	} else {
		fmt.Println(y, "é maior que", x)

	}

}
