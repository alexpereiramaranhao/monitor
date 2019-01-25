package main

import (
	"fmt"
)

func main() {

	var opcao int

	montarMenu()

	fmt.Scan(&opcao)

	if opcao == 1 {
		fmt.Println("Monitoramento")
	} else if opcao == 2 {
		fmt.Println("Logs")
	} else if opcao == 0 {
		fmt.Println("Saindo...")
	} else {
		fmt.Println("Opção inválida:")
		montarMenu()
	}

}

func montarMenu() {
	versao := 1.1

	fmt.Println("Monitor sistema, versão", versao)

	fmt.Println("Menu")

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}
