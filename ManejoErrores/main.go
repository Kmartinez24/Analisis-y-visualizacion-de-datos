package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func cuadradoPar2(n int) (int, bool) {
	if n%2 == 0 {
		return n * n, true
	} else {
		return 0, false
	}
}

func cuadradoPar(n int) (int, error) {
	if n%2 == 0 {
		return n * n, nil
	} else {
		return 0, errors.New("Tiene que ser par")
	}
}

func main() {
	numero, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("El error es: ", err)
	} else {
		cuad, err := cuadradoPar(numero)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("El cuadrado es: ", cuad)
		}
	}
}
