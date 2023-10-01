package main

import "fmt"

// Faça um algoritmo que leia o salário de um funcionário e calcule o seu novo
// salário com um aumento de 10% se o salário for menor que
// R$ 1000,00; ou de 5% se o salário for maior ou igual a R$ 1000,00.
func main() {
	var sal float64
	fmt.Println("bom dia, qual é o seu salario?")
	fmt.Scan(&sal)
	if sal < 1000 {
		fmt.Println("seu salario com um aumento de 10% é", sal*1.1)

	} else {
		fmt.Println("seu salario com um aumento de 5% é", sal*1.05)
	}

}
