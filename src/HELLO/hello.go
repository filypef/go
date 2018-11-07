package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitory = 3 //number of times monitoring will run
const delay = 2    //number of delay

func main() {
	var name string
	fmt.Print("Por favor informar seu nome: ")
	fmt.Scan(&name)

	for {
		helloName(name)

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

func helloName(name string) {
	if name == "" {
		fmt.Print("Por favor informar seu nome: ")
		fmt.Scan(&name)
	}
	fmt.Println("****Seja bem-vindo ", name, " ****")
}

func displayMenu() {

	fmt.Println("--------------------------------------")
	fmt.Println("Selecione uma das opções abaixo:")
	fmt.Println("1. Iniciar o monitoramento")
	fmt.Println("2. Exibir os logs")
	fmt.Println("0. Sair")
	fmt.Println("--------------------------------------")
}

func readCommand() int {
	var comando int
	fmt.Scan(&comando)

	return comando
}

func exitSucess() {
	os.Exit(0)
}

func initMonitoring(command int) {
	// fmt.Println("Monitoramento...", command)
	// // site := "https://random-status-code.herokuapp.com/"

	sites := []string{"http://www.alura.com.br", "http://www.globo.com", "http://www.natura.com.br/404"}

	//for tradicional
	// for i := 0; i < len(sites); i++ {
	// 	fmt.Println(sites[i])
	// }

	for i := 0; i < monitory; i++ {
		//FOR estilo for each
		for i, site := range sites {
			testSite(i, site)
		}
		fmt.Println("")
		fmt.Println(i)
		time.Sleep(delay * time.Minute)
	}
}

//function test site
func testSite(i int, site string) {
	response, _ := http.Get(site)
	if response.StatusCode == 200 {
		fmt.Println("Testando site: ", site, " foi carregado com sucesso")
	} else {
		fmt.Println("Testando site: ", site, " o site apresentou erro")
	}
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
