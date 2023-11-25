package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	exibeIntroducao()
	for {
		exibeMenu()
		comando := leComando()

		switch comando {
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs")
		default:
			fmt.Println("Não existe este comando")
			os.Exit(-1)
		}
	}
}

// func devolveNomeEIdade() (string, int) {
// 	nome := "Iago"
// 	idade := 23
// 	return nome, idade
// }

func exibeIntroducao() {
	nome := "Iago"
	versao := 0.1
	fmt.Println("Olá, sr.", nome, "sua idade é")
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)
	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando..")
	site := "https://www.alura.com.br"
	// site := "https://random-status-code.herokuapp.com/"
	resp, _ := http.Get(site)
	fmt.Println(resp.StatusCode)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site", site, "esta com problema. Status code:", resp.StatusCode)
	}

}
