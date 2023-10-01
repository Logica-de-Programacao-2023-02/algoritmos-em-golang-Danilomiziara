package main

import "fmt"

//Faça um algoritmo que leia três números inteiros e mostre a soma entre eles.

func main() {
	var x int
	var y int
	var z int
	var resultado int
	fmt.Println("fale 3 numeros inteiros")
	fmt.Scan(&x, &y, &z)
	resultado = x + y + z
	fmt.Println("a soma dos 3 equivale a:", resultado)
	fmt.Println("CHUPA KAIKE EU CONSEGUI!!!!!")

}
