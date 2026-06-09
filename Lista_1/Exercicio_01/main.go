package main

import (
	"fmt"
)

/*Primeira questão da lista*/
/*Exercício 1.
Considere o slice []int{4, 17, 2, 9, 31, 6, 14}. Percorra todos os elementos e imprima cada um em uma linha separada.*/

var list = []int{4, 17, 2, 9, 31, 6, 14}

func main() {
	for i, value := range list {
		fmt.Println("Posição:", i, "Valor:", value)
	}
}
