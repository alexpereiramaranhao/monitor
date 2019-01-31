package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	versao := 1.1

	fmt.Println("Monitor sistema, versão", versao)

	for {
		montarMenu()

		opcao := lerComando()

		switch opcao {
		case 0:
			fmt.Println("Saindo...")
			os.Exit(0)
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Logs")
		default:
			fmt.Println("Opção inválida:")
			os.Exit(-1)
		}
	}
}

func lerComando() int {
	var opcao int

	fmt.Scan(&opcao)

	return opcao
}

func montarMenu() {

	fmt.Println("Menu")

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func iniciarMonitoramento() {
	site := "https://random-status-code.herokuapp.com/"

	response, _ := http.Get(site)

	statusCode := response.StatusCode

	if statusCode == 200 {
		fmt.Println("Status ok:", statusCode)
	} else {
		fmt.Println("Erro:", statusCode)
	}
}
