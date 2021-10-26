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

const timesTest = 5
const delay = 10

func main() {

	fmt.Println("Olá, o que você deseja?")
	sites := readFile()
	for {
		showMenu()
		option := readOption()
		switch option {
		case 1:
			monitoringSites(sites)
		case 2:
			getLogs()
		case 0:
			fmt.Println("Adeus")
			os.Exit(0)
		default:
			fmt.Println("Essa opção não existe")
			fmt.Println("")
		}
	}

}

func showMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir os logs")
	fmt.Println("0 - Sair")
}

func monitoringSites(sites []string) {
	fmt.Println("Monitorando")

	for i := 0; i < timesTest; i++ {
		for _, site := range sites {
			testSite(site)
		}
		fmt.Println("")
		time.Sleep(delay * time.Second)
	}
	fmt.Println("")
}

func getLogs() {
	file, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Houve o erro", err)
		return
	}

	fmt.Println(string(file))
}

func readOption() int {
	var comando int
	fmt.Scan(&comando)
	return comando
}

func testSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("O site", site, "está rodando!")
		generateLog(site, true)
	} else {
		fmt.Println("O site", site, "não está rodando, erro ", resp.StatusCode)
		generateLog(site, false)
	}
}

func readFile() []string {
	file, err := os.Open("sites.txt")
	reader := bufio.NewReader(file)
	sites := []string{}

	if err != nil {
		fmt.Println("Ocorreu o erro:", err)
		return nil
	}

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		sites = append(sites, strings.TrimSpace(line))
	}
	file.Close()
	return sites
}

func generateLog(site string, status bool) {

	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "-ONLINE:" + strconv.FormatBool(status) + "\n")
	file.Close()
}
