package main

import "fmt"

func main() {

	list := []int{38, 27, 43, 3, 9, 82, 10}

	for i := 0; i < len(list)-1; i++ {
		for j := 0; j < len(list)-1-i; j++ {
			if list[j] < list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}

	fmt.Println(list)
}
