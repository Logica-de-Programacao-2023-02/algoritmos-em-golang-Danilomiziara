package main

import "fmt"

//Faça um algoritmo que leia um número inteiro e mostre o seu sucessor e antecessor.

func main() {
	var num int
	fmt.Println("fala um numero inteiro")
	fmt.Scan(&num)
	fmt.Println("o antecessor do seu número é", num-1, "e o sucessor é", num+1)

}
