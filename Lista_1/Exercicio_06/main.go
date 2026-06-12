package main

import "fmt"

var frutas []string = []string{"Banana", "maça", "uva", "laranja", "abacaxi"}

func main() {
	for i, fruta := range frutas {
		if fruta == "uva" {
			fmt.Println("Posição uva:",i)
		}else{
			fmt.Println("Uva não existe nessa lista...")
			break
		}
	}
}