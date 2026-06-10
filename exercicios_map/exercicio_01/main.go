package main

import "fmt"

var pessoas []map[string]string

func main() {
	var opcao int

	for {
		fmt.Println("----------MENU----------")
		fmt.Println("1 - Cadastrar Pessoa")
		fmt.Println("2 - Listar Pessoa")
		fmt.Println("3 - Sair")

		fmt.Print("Escolha uma opção:")
		fmt.Scan(&opcao)

		if opcao == 1 {

			adicionarPessoa()

		} else if opcao == 2 {

			listarPessoas()

		} else if opcao == 3 {
			fmt.Println("Saindo...")
			break
		} else {
			fmt.Println("Opção inválida!!!")
		}
	}
}

func adicionarPessoa() {

	var nome string
	var sexo string

	fmt.Print("Nome:")
	fmt.Scan(&nome)
	fmt.Print("Sexo:")
	fmt.Scan(&sexo)

	pessoas = append(pessoas, map[string]string{
		"nome": nome,
		"sexo": sexo,
	})

	fmt.Println("Cadastro realizado com sucesso!!")

}

func listarPessoas() {
	for _, pessoa := range pessoas {
		fmt.Println("Nome: ", pessoa["nome"])
		fmt.Println("Sexo: ", pessoa["sexo"])
	}
}
