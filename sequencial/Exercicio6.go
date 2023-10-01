package main

import "fmt"

//Faça um algoritmo que leia o salário de um funcionário
//e calcule o seu novo salário com um aumento de 15%.

func main() {
	var sal float64
	fmt.Println("olá, qual o seu salario?")
	fmt.Scan(&sal)
	fmt.Println("com um aumento de 15%, seu salario será:", sal*1.15)

}
