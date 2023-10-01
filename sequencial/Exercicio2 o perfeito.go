package main

import (
	"fmt"
)

//Faça um algoritmo que leia um número inteiro e mostre o seu dobro, triplo e quadruplo

func main() {
	var x int
	var res1 int
	var res2 int
	var res3 int
	var resposta string
	fmt.Println("diga um numero inteiro")
	fmt.Scan(&x)
	res1 = x * 2
	res2 = x * 3
	res3 = x * 4
	fmt.Println("Você quer saber o dobro? s/n")
	fmt.Scan(&resposta)

	if resposta == "s" {
		fmt.Println("o dobro do seu numero é:", res1)
	} else if resposta == "n" {
		fmt.Println("tabo")

	} else {
		fmt.Println("vai toma no cu")
	}

	fmt.Println("Você quer saber o triplo? s/n:")
	fmt.Scan(&resposta)

	if resposta == "s" {
		fmt.Println("o triplo do seu numero é:", res2)
	} else if resposta == "n" {
		fmt.Println("tabo")

	} else {
		fmt.Println("vai toma no cu")
	}

	fmt.Println("Você quer saber o quadruplo? s/n:")
	fmt.Scan(&resposta)

	if resposta == "s" {
		fmt.Println("o quadruplo do seu numero é:", res3)
	} else if resposta == "n" {
		fmt.Println("tabo")

	} else {
		fmt.Println("vai toma no cu")
	}
}
