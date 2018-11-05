package main

import "fmt"

func main() {

	fmt.Println("Selecione uma das opções abaixo:")

	fmt.Println("1. Iniciar o monitoramento")

	fmt.Println("2. Exibir os logs")

	fmt.Println("0. Sair")

	var comando int
	fmt.Scan(&comando)

	fmt.Println("O comando é: ", comando)

	// if comando == 1 {
	// 	fmt.Println("Monitoramento...", comando)
	// } else if comando == 2 {
	// 	fmt.Println("Gerando logs...", comando)
	// } else if comando == 0 {
	// 	fmt.Println("Saindo...", comando)
	// } else {
	// 	fmt.Println("Opção não reconhecida")
	// }

	switch comando {
	case 1:
		fmt.Println("Monitoramento...", comando)
	case 2:
		fmt.Println("Gerando logs...", comando)
	case 0:
		fmt.Println("Saindo...", comando)
	default:
		fmt.Println("Opção não reconhecida")
	}
}
