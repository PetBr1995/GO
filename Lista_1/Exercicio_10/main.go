package main

import "fmt"

var slice []int = []int{100,200,300,400,500}

func main() {
	fmt.Println("Teste")

	slice = append(slice[:2],slice[2+1:]... )
	fmt.Println(slice)
}