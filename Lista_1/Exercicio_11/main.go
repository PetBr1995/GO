package main

import "fmt"

var slice []int = []int{14, 22, 8, 35, 19, 42, 7, 28, 13, 31}
var soma float32
var media float32

func main() {
	for _, value := range slice {
		soma += float32(value)
	}

	media = soma / float32(len(slice))
	fmt.Println("Média aritmética do slice:",media)
}