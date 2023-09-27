package main

import (
	"fmt"
	"strconv"
)

var (
	nome string
	n1   int
	n2   int32
)

func Hello(nome string) {
	fmt.Println("Hello", nome)
}

func sum(a, b int) int {
	return a + b
}

func convertAndSum(a int, b string) (total int, err error) {

	i, err := strconv.Atoi(b)

	if err != nil {
		return
	}
	total = a + i

	return
}

var x, y = 10, 20

func main() {

	hello("Gustavo Barros")

	fmt.Println("total:", sum(10, 5))
	total, err := convertAndSum(10, "23")
	fmt.Println("total:", total, err)

}
