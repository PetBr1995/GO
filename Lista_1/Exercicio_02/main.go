package main

import (
	"fmt"
)

/*Segunda questão da lista*/
/*Exercício 2.
Dado o slice []int{50, 23, 7, 88, 15, 42, 3, 66}, encontre e imprima o maior e o menor valor presentes.*/

var list = []int{50, 23, 7, 88, 15, 42, 3, 66}
var maiorNumero int = list[0]
var menorNumero int = list[0]

func main() {
	for _, value := range list {
		if value >= maiorNumero {
			maiorNumero = value
		}
		if value <= menorNumero {
			menorNumero = value
		}
	}

	fmt.Println("O maior número da lista é: ", maiorNumero)
	fmt.Println("O menor número da lista é: ", menorNumero)
}
