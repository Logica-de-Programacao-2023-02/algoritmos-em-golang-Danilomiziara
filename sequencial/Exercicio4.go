package main

import "fmt"

//Faça um algoritmo que leia três números reais e mostre a média ponderada entre eles,
//com pesos 2, 3 e 5, respectivamente.

func main() {
	var x int
	var y int
	var z int

	fmt.Println("fale 3 numeros inteiros")
	fmt.Scan(&x, &y, &z)
	fmt.Println("a media ponderada é:", float64((x*2+y*3+z*5))/(2+3+5))

}
