package main

import "fmt"

// Faça um algoritmo que leia a idade de uma pessoa e mostre a sua classificação,
// de acordo com a seguinte tabela:
// até 9 anos: mirim
// 10 a 13 anos: infantil
// 14 a 17 anos: juvenil
// maiores de 18 anos: adulto
func main() {
	var idade int
	fmt.Println("qual a sua idade?")
	fmt.Scan(&idade)
	if idade <= 9 {
		fmt.Println("mirim")

	} else if idade > 9 && idade <= 13 {
		fmt.Println("infantil")

	} else if idade > 13 && idade <= 17 {
		fmt.Println("juvenil")
	} else {
		fmt.Println("adulto!!!!")
	}
}
