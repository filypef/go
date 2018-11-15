package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitory = 2 //number of times monitoring will run
const delay = 2    //number of delay

func main() {

	var name string
	fmt.Print("Por favor informar seu nome: ")
	fmt.Scan(&name)

	readSiteOfFile()

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

	sites := readSiteOfFile()

	//for tradicional
	// for i := 0; i < len(sites); i++ {
	// 	fmt.Println(sites[i])
	// }

	for i := 0; i < monitory; i++ {
		//FOR estilo for eachs
		for _, site := range sites {
			testSite(site)
		}
		fmt.Println("")
		time.Sleep(delay * time.Minute)
	}
}

//function test site
func testSite(site string) {
	response, err := http.Get(site)

	if err != nil {
		fmt.Println("Errou")
	}

	if response.StatusCode == 200 {
		fmt.Println("Testando site: ", site, " foi carregado com sucesso")
		createLog(site, true)
	} else {
		fmt.Println("Testando site: ", site, " o site apresentou erro")
		createLog(site, false)
	}
}

func readSiteOfFile() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		println("Ocorreu um erro readSite: ", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}
	arquivo.Close()

	return sites
}

func createLog(site string, statusCode bool) {

	arquivo, err := os.OpenFile("log.text", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Print(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- online: " + strconv.FormatBool(statusCode) + "\n")

	arquivo.Close()
}

// Exercicio de slice no GO
// func exibeNomes() {
// 	nomes := []string{"Filype", "Gabriela", "Marli", "Danielly"}

// 	fmt.Println("Os nomes são", nomes)

// 	fmt.Println("Tamanho do slice: ", len(nomes))
// 	fmt.Println("Capacidade do meu array: ", cap(nomes))
// 	fmt.Println("--------------------------------------")

// 	nomes = append(nomes, "Josuel")

/// 	fmt.Println("Tamanho do slice: ", len(nomes))
// 	fmt.Println("Capacidade do meu array: ", cap(nomes))

// 	fmt.Println("--------------------------------------")

// 	cartas := []int{1, 2, 3, 5, 8, 13, 21}
// 	fmt.Println("Quantidade de cartas necessárias ", len(cartas))
// }
