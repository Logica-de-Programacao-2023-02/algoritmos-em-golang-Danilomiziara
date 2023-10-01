package main

import "fmt"

//Faça um algoritmo que leia o preço de um produto e mostre o seu valor com desconto de 10%.

func main() {
	var preço float64
	fmt.Println("Qual o valor original do produto?")
	fmt.Scan(&preço)
	fmt.Println("O seu produto está com 10% de desconto, seu valor no momento é", preço*0.9)

}
