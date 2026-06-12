package main

import "fmt"

var slice []int = []int{1,2,3,4,5,6,7,8,9,10}
var slicePares []int 

func main() {
	for _,numero := range slice{
		if numero % 2 == 0 {
			slicePares = append(slicePares, numero)
		}
	}
	fmt.Println(slicePares)
}