package main

import "fmt"

//Faça um algoritmo que leia o número de dias trabalhados por um
//funcionário e o valor da sua diária e calcule o seu salário.

func main() {
	var trab int
	var dia int
	fmt.Println("olá quantos dias por mês você trabalha?")
	fmt.Scan(&trab)
	fmt.Println("qual é o valor da sua diária?")
	fmt.Scan(&dia)
	fmt.Println("seu salário mensal será:", trab*dia)

}
