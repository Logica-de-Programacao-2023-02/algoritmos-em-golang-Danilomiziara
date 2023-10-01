package main

import "fmt"

// Faça um algoritmo que leia um número inteiro entre 1 e 7 e mostre o dia da
// semana correspondente (1 = domingo, 2 = segunda-feira, etc.).
func main() {
	var x int
	fmt.Println("digite um número de um a sete")
	fmt.Scan(&x)
	if x <= 0 {
		fmt.Println("inválido")

	} else if x == 1 {
		fmt.Println("domingo")

	} else if x == 2 {
		fmt.Println("segunda")

	} else if x == 3 {
		fmt.Println("terça")
	} else if x == 4 {
		fmt.Println("quarta")
	} else if x == 5 {
		fmt.Println("quinta")
	} else if x == 6 {
		fmt.Println("sexta")
	} else if x == 7 {
		fmt.Println("sabado")
	}

}
