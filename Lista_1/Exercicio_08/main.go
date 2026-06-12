package main

import "fmt"

var slice []int = []int{12,45,7,23,89,34,67,11,90}

func main() {
	for i,numero := range slice{
		if numero == 56{
			fmt.Println("Posição no número ",numero,":",i)
		}else{
			fmt.Println("O número não exite na lista...")
			break
		}
	}
}