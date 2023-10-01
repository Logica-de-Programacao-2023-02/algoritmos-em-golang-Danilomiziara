package main

import "fmt"

//Faça um algoritmo que leia o peso de uma pessoa em quilos e converta para libras.

func main() {
	var peso float64
	fmt.Println("olá, qual o seu peso?")
	fmt.Scan(&peso)
	fmt.Println("seu peso em libras é:", peso*2.205, "lbs")

}
