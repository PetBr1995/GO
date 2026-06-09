package main

import "fmt"

func main() {
	var soma float64
	numeros := [5]float64{2, 5, 22, 5, 78}

	for _, value := range numeros {
		soma += value
		fmt.Print(value, ",")
	}

	fmt.Print("O valor da soma dos números é: ", soma)
}
