package main

import (
	"fmt"
	"os"
)

func main() {

	displayMenu()

	command := readCommand()

	switch command {
	case 1:
		fmt.Println("Monitoramento...", command)
	case 2:
		fmt.Println("Gerando logs...", command)
	case 0:
		fmt.Println("Saindo...", command)
		exitSucess()
	default:
		fmt.Println("Opção não reconhecida")
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

	fmt.Println("O comando é: ", comando)
	return comando
}

func exitSucess() {
	os.Exit(0)
}
