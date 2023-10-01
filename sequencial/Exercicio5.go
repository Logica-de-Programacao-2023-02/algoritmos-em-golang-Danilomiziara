package main

import "fmt"

//Faça um algoritmo que leia a idade de uma pessoa em anos e mostre a idade em dias.

func main() {
	var idade int
	fmt.Println("diga sua idade em anos")
	fmt.Scan(&idade)
	fmt.Println("sua idade em dias é:", idade*365)

}
