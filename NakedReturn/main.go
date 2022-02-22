package main

import "fmt"

func sumaCuadrados(n int) (suma int) {
	if n >= 5 {
		for i := 0; i < n; i++ {
			suma = suma + (n * n)
		}
	} else {
		n = 0
	}
	return
}

func cuadrado(n int) (cuadrado int) {
	cuadrado = n * n
	return
}

func main() {
	fmt.Println(sumaCuadrados(5))
	fmt.Println(cuadrado(5))
}
