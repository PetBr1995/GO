package main

import "fmt"

var slice []int = []int{5, 3, 8, 1, 9, 2, 7, 4, 6}

func main() {

	for i := 0; i < len(slice)-1; i++ {
		menor := i

		for j := i + 1; j < len(slice); j++ {
			if slice[j] < slice[menor] {
				menor = j
			}
		}
		slice[i], slice[menor] = slice[menor], slice[i]
	}
	fmt.Println(slice)
}