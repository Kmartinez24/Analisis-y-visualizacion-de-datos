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
		n1, error1 := strconv.Atoi(os.Args[1])
		if error1 != nil {
			fmt.Println("Error en el primer argumento")
			return
		} else {
			n2, error2 := strconv.Atoi(os.Args[2])
			if error2 != nil {
				fmt.Println("Error en el segundo argumento")
				return

			} else {
				fmt.Printf("El resultado de la operacion %d / %d es: %2.2f", n1, n2, dividir(n1, n2))
			}
		}
	}
}
