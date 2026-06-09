package main

import "fmt"

type Produto struct {
	Nome  string
	Preco float64
}

var produtos []Produto

func main() {

	var opcao int
	for {

		fmt.Println("1 - Cadastrar produto")
		fmt.Println("2 - Listar produtos")
		fmt.Println("3 - Mostrar quantidade de produtos cadastrados")
		fmt.Println("0 - Sair")
		fmt.Println("Escolha uma opção:")

		fmt.Scan(&opcao)

		if opcao == 1 {

			cadastrarProduto()

		} else if opcao == 2 {

			listarProdutos()

		} else if opcao == 3 {
			fmt.Println("Opção 3")
		} else if opcao == 0 {
			fmt.Println("Saindo...")
			break
		} else {
			fmt.Println("Opção inválida!!")
		}

	}

}

func cadastrarProduto() {
	var nome string
	var preco float64

	fmt.Println("Digite o nome do produto: ")
	fmt.Scan(&nome)
	fmt.Println("Digite o valor do produto: ")
	fmt.Scan(&preco)

	produtos = append(produtos, Produto{
		Nome:  nome,
		Preco: preco,
	})

	fmt.Println("Produto cadastrado com sucesso!!!")
}

func listarProdutos() {
	for _, produto := range produtos {
		fmt.Println(produto.Nome)
		fmt.Println(produto.Preco)
	}
}
