package main

import "fmt"

/*Exercício 3. Considere o slice []int{10, 20, 30, 40, 50, 60, 70, 80, 90}. Inverta a ordem dos elementos sem usar um slice auxiliar e imprima o resultado.*/

func main() {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(list)
	for i := 0; i < len(list)/2; i++ {
		j := len(list) - 1 - i

		list[i], list[j] = list[j], list[i]
	}
	fmt.Println(list)

}
