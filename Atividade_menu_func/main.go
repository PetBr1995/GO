package main

import "fmt"

type Pessoa struct {
	Nome string
	Sexo string
}

var pessoas []Pessoa

func main() {

	var opcao int
	for opcao != 4 {

		fmt.Println("1 - Adicionar nome e sexo de uma pessoa")
		fmt.Println("2 - Listar todos os nomes")
		fmt.Println("3 - Total de homens e total de mulheres")
		fmt.Println("4 - Sair")
		fmt.Println("Escolha uma opção: ")

		fmt.Scan(&opcao)

		if opcao == 1 {

			adicionarPessoa()

		} else if opcao == 2 {

			listarPessoas()

		} else if opcao == 3 {

			qtdHomensEMulheres()

		} else if opcao == 4 {
			fmt.Println("Saindo...")
		} else {
			fmt.Println("Opção inválida")
		}
	}

}

func adicionarPessoa() {
	var nome string
	var sexo string

	fmt.Print("Digite o nome da pessoa:")
	fmt.Scan(&nome)
	fmt.Print("Digite o sexo da pessoa: Masculino(M) / Feminino(F)")
	fmt.Scan(&sexo)

	pessoas = append(pessoas, Pessoa{
		Nome: nome,
		Sexo: sexo,
	})
}

func listarPessoas() {
	for _, pessoa := range pessoas {
		fmt.Println("Nome: ", pessoa.Nome)
	}
}

func qtdHomensEMulheres() {
	total_homens := 0
	total_mulheres := 0

	fmt.Println("------Total de homens e mulheres------")

	for _, pessoa := range pessoas {
		if pessoa.Sexo == "M" || pessoa.Sexo == "m" {
			total_homens += 1
		} else if pessoa.Sexo == "F" || pessoa.Sexo == "f" {
			total_mulheres += 1
		}
	}

	fmt.Println("Total de Homens: ", total_homens)
	fmt.Println("Total de Mulheres: ", total_mulheres)

}
