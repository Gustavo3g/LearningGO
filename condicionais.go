package main

import (
	"fmt"
	"log"
	"os"
)

var (
	cara, coroa int
)

func jogarMoeda(lado string) {

	switch {
	case lado == "cara":
		fmt.Println("Caiu Cara")
		cara++

	case lado == "coroa":
		fmt.Println("Caiu Coroa")
		coroa++

	default:
		fmt.Println("Caiu em pé")
	}
}

func main() {

	a, b := 200, 50

	if a > b {
		fmt.Println("A é maior que B")
	} else if b > a {
		fmt.Println("B é maior que A")
	} else {
		fmt.Println("A e B são iguais")
	}

	switch {
	case a > b:
		fmt.Println("A é maior que B")
	case b > a:
		fmt.Println("B é maior que A")
	default:
		fmt.Println("A e B são iguais")

	}

	file, err := os.Open("hello.txt")

	if err != nil {
		log.Panic(err)
	}

	data := make([]byte, 100)

	if _, err := file.Read(data); err != nil {
		log.Panic(err)
	}

	fmt.Println(string(data))

}
