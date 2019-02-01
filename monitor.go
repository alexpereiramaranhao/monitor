package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const MONITORAMENTOS = 5
const DELAY = 2

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
	sites := lerArquivo()

	for i := 0; i < MONITORAMENTOS; i++ {
		for _, site := range sites {
			testarSite(site)
		}

		time.Sleep(DELAY * time.Second)
	}

}

func testarSite(site string) {
	response, err := http.Get(site)

	if err != nil {
		fmt.Println("Erro ao testar o site. ", err)

		return
	}

	statusCode := response.StatusCode

	if statusCode == 200 {
		fmt.Println(site, "- status ok:", statusCode)
	} else {
		fmt.Println(site, " - erro:", statusCode)
	}
}

func lerArquivo() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro ao ler o arquivo:", err)
		return nil
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		if err == io.EOF {
			break
		}

		sites = append(sites, linha)
	}

	arquivo.Close()

	fmt.Println(sites)

	return sites
}
