package main

import (
	"fmt"
	"os"
)

func main() {

	for {
		displayMenu()

		command := readCommand()

		switch command {
		case 1:
			initMonitoring(command)
		case 2:
			fmt.Println("Gerando logs...", command)
		case 0:
			fmt.Println("Saindo...", command)
			exitSucess()
		default:
			fmt.Println("Opção não reconhecida")
		}
	}
}

func displayMenu() {
	fmt.Println("Selecione uma das opções abaixo:")

	fmt.Println("1. Iniciar o monitoramento")

	fmt.Println("2. Exibir os logs")

	fmt.Println("0. Sair")
}

func readCommand() int {
	var comando int
	fmt.Scan(&comando)

	// fmt.Println("O comando é: ", comando)

	return comando
}

func exitSucess() {
	os.Exit(0)
}

func initMonitoring(command int) {
	// fmt.Println("Monitoramento...", command)
	// // site := "https://random-status-code.herokuapp.com/"

	sites := []string{"alura.com.br", "globo.com", "natura.com.br"}

	//for tradicional
	// for i := 0; i < len(sites); i++ {
	// 	fmt.Println(sites[i])
	// }

	for _, site := range sites {
		fmt.Println(site)
	}

	// response, _ := http.Get(site)

	// if response.StatusCode == 200 {
	// 	fmt.Println("O site: ", site, "foi carregado com sucesso", response.StatusCode)
	// } else {
	// 	fmt.Print("O site: ", site, " está com erro ", response.StatusCode)
	// }
}

// Exercicio de slice no GO
// func exibeNomes() {
// 	nomes := []string{"Filype", "Gabriela", "Marli", "Danielly"}

// 	fmt.Println("Os nomes são", nomes)

// 	fmt.Println("Tamanho do slice: ", len(nomes))
// 	fmt.Println("Capacidade do meu array: ", cap(nomes))
// 	fmt.Println("--------------------------------------")

// 	nomes = append(nomes, "Josuel")

// 	fmt.Println("Tamanho do slice: ", len(nomes))
// 	fmt.Println("Capacidade do meu array: ", cap(nomes))

// 	fmt.Println("--------------------------------------")

// 	cartas := []int{1, 2, 3, 5, 8, 13, 21}
// 	fmt.Println("Quantidade de cartas necessárias ", len(cartas))
// }
