package main

import (
	"io/ioutil"
	"fmt"
	"os"
	"net/http"
	"time"
	"bufio"
	"io"
	"strings"
	"strconv"
)

const monitoramentos = 3
const delay = 5

func main() {
	exibeIntro()
	for {
		exibeMenu()
		comando := leComando()
	
		switch comando {
		case 0: 
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		case 1: 
			iniciarMonitoramento()
		case 2: 
			imprimeLogs()
		default: 
			fmt.Println("Comando desconhecido")
			os.Exit(-1)
		}
	}

}

func exibeIntro() {
	name := "Rangel"
	version := 1.1
	
	fmt.Println("Olá Sr.", name)
	fmt.Println("Versão: ", version)
}

func exibeMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")
}

func leComando() int {
	var comando int 
	fmt.Scan(&comando)

	fmt.Println("O comando escolhido foi:", comando)
	return comando
}

func iniciarMonitoramento() {
	fmt.Println("Monitoramento...")
	var sites = leSitesDoArquivo()

	for index := 0; index < monitoramentos; index++ {
		for _, site := range sites {
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
	fmt.Println("")
}

func testaSite(site string) {
	resp, err := http.Get(site)	

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
		
	if resp.StatusCode == 200 {
		fmt.Println("Site: ", site, "carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site: ", site, " está com problemas. Status Code: ", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
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

func registraLog (site string, status bool){
	arquivo, err := os.OpenFile("log.txt", os.O_CREATE | os.O_RDWR | os.O_APPEND, 0666)
	
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(arquivo)
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")
	arquivo.Close()
}

func imprimeLogs(){
	fmt.Println("Exibindo logs...")
	arquivo, err := ioutil.ReadFile("log.txt")
	
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(arquivo))
}