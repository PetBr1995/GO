package main

import "fmt"

var produtos = make(map[string]float32)

func main() {
	var opt int

	for {
		fmt.Println("----------MENU----------")
		fmt.Println("1 - Cadastrar produto")
		fmt.Println("2 - Listar produtos")
		fmt.Println("0 - Sair")

		fmt.Println("Escolha uma opção:")
		fmt.Scan(&opt)

		if opt == 1 {
			cadastrarProduto()
		} else if opt == 2 {
			listarProdutos()
		} else if opt == 0 {
			fmt.Println("Saindo...")
			break
		} else {
			fmt.Println("Opção inválida...")
			continue
		}

	}
}

func cadastrarProduto() {
	var nome string
	var preco float32

	fmt.Println("Digite o nome do produto:")
	fmt.Scan(&nome)
	fmt.Println("Digite o preço do produto:")
	fmt.Scan(&preco)

	produtos[nome] = preco

	fmt.Println("Produto cadastrado com sucesso!")
}

func listarProdutos() {
	for nome, preco := range produtos {
		fmt.Println(" --------------------------------------")
		fmt.Println("|Produto:", nome, "| Preço:R$", preco, "|")
		fmt.Println(" --------------------------------------")
	}
}
