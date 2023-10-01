package main

import "fmt"

// Faça um algoritmo que leia vários números inteiros e mostre a média aritmética entre eles.
// A leitura deve ser interrompida quando for lido o valor zero.
func main() {
	var x, y, z int
	fmt.Println("digite varios numeros, caso queira parar digite 0")
	for {
		fmt.Scan(&y)
		if y == 0 {
			break
		}
		x += y
		z++
	}
	if z == 0 {
		fmt.Println("nenhum numero foi inserido")

	} else {
		media := x / z
		fmt.Println("a media é", media)

	}

}
