package main

import (
	"fmt"
	"os"
	"strconv"
)

func dividir(n1, n2 int) float64 {
	return float64(n1) / float64(n2)
}

func main() {

	if len(os.Args) < 3 {
		fmt.Println("El formato es go run main.go n1 n2")
	} else {
		n1str := os.Args[1]
		n2str := os.Args[2]
		n1, _ := strconv.Atoi(n1str)
		n2, _ := strconv.Atoi(n2str)
		res := dividir(n1, n2)

		fmt.Printf("El resultado de la operacion %s / %s es: %2.2f", n1str, n2str, res)
	}
}
