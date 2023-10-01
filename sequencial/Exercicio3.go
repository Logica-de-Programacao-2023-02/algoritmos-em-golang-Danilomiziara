package main

import "fmt"

//Faça um algoritmo que leia o peso e a altura de uma pessoa e calcule o seu IMC
//(Índice de Massa Corporal).

func main() {
	var altura float64
	var peso float64
	var name string
	var res float64
	var imc float64
	fmt.Println("olá qual é seu nome?")
	fmt.Scan(&name)
	fmt.Println(name, "qual é seu peso?")
	fmt.Scan(&peso)
	fmt.Println(name, "qual a sua altura?")
	fmt.Scan(&altura)
	res = altura * altura
	imc = peso / res
	fmt.Println(name, "o seu imc é:", imc)

}
