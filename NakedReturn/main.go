package main

import (
	"fmt"
	"os"
	"strconv"
)

func sumaCuadrados(n int) (suma int) {
	for i := 0; i <= n; i++ {
		if i%2 == 0 {
			suma = suma + (i * i)
		}
	}
	return
}

func sumaCubos(n int) (suma int) {
	for i := 0; i <= n; i++ {
		if i%2 != 0 {
			suma = suma + (i * i * i)
		}
	}
	return
}

func parImpar(n int) int {
	if n%2 != 0 {
		return sumaCubos(n)
	} else {
		return sumaCuadrados(n)
	}

}

func main() {
	Numero, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Hay un error en el codigo")
	} else {
		fmt.Println("El resultado es", parImpar(Numero))
	}
}
