package main

import "fmt"

var lista []int = []int{1, 3, 5, 7, 9}
var lista2 []int = []int{2, 4, 6, 8, 10}

func main() {
	lista = append(lista, lista2...)

	for i := 0; i < len(lista)-1; i++ {
		for j := 0; j < len(lista)-1-i; j++ {
			if lista[j] > lista[j+1] {
				lista[j], lista[j+1] = lista[j+1], lista[j]
			}
		}
	}

	fmt.Println(lista)
}