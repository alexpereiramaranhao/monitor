package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// Controle do monitoramento
const (
	MONITORAMENTOS = 5

	DELAY = 2
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
			imprimirLogs()
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
		escreverLog(site, true)
	} else {
		fmt.Println(site, " - erro:", statusCode)
		escreverLog(site, false)
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

func escreverLog(site string, status bool) {
	arquivo, err := os.OpenFile("status.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Erro ao escrever no arquivo.", arquivo)

		return
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimirLogs() {
	arquivo, err := ioutil.ReadFile("status.log")

	if err != nil {
		fmt.Println("Erro ao ler arquivo", arquivo)

		return
	}

	fmt.Println(string(arquivo))
}
