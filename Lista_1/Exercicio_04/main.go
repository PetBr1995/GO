package main

import "fmt"

func main() {

	var qtdN5 int

	list := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}

	for _, value := range list {

		if value == 5 {
			qtdN5 += 1
		}
	}

	fmt.Println(list)
	fmt.Println("Quantidade de numeros 5:", qtdN5)
}
