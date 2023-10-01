package main

import "fmt"

// Faça um algoritmo que leia vários números inteiros e mostre o maior deles.
// A leitura deve ser interrompida quando for lido o valor zero.
func main() {
	var x, y int
	fmt.Println("digite varios numeros, caso queira parar digite 0")
	for {
		fmt.Scan(&y)
		if y == 0 {
			fmt.Println(x, "é o maior")
			break
		}
		if y > x {
			x = y
		}

	}

}
