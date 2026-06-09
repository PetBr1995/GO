package main

import (
	"fmt"
)

/*Primeira questão da lista*/

var list = []int{4, 17, 2, 9, 31, 6, 14}

func main() {
	for i, value := range list {
		fmt.Println("Posição:", i, "Valor:", value)
	}
}
