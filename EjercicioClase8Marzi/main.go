package main

import (
	"fmt"
	"math/rand"
	"time"
)

func paresImparesSlice(nums []int, bandera string) (sum int) {
	if bandera == "Pares" || bandera == "pares" || bandera == "2" {
		fmt.Println("Pares")
		for _, num := range nums {
			if num%2 == 0 {
				sum += num
			}
		}
	} else if bandera == "Impares" || bandera == "impares" || bandera == "1" {
		fmt.Println("Impares")
		for _, num := range nums {
			if num%2 == 1 {
				sum += num
			}
		}
	}
	return
}

func paresImparesVarica(bandera string, nums ...int) (sum int) {
	if bandera == "Pares" || bandera == "pares" || bandera == "2" {
		fmt.Println("Pares")
		for _, num := range nums {
			if num%2 == 0 {
				sum += num
			}
		}
	} else if bandera == "Impares" || bandera == "impares" || bandera == "1" {
		fmt.Println("Impares")
		for _, num := range nums {
			if num%2 == 1 {
				sum += num
			}
		}
	}
	return
}

func main() {
	rand.Seed(time.Now().UnixNano())
	numeros := make([]int, 10)
	for i := 0; i < 10; i++ {
		numeros[i] = rand.Intn(10)
	}
	fmt.Println(numeros)
	fmt.Println("La suma de pares con slice es: ", paresImparesSlice(numeros, "2"))
	fmt.Println("La ssma de impares con slice es: ", paresImparesSlice(numeros, "1"))
	fmt.Println("La suma de pares con multiargumento es: ", paresImparesVarica("2", 5, 6, 7, 2, 3, 1, 7, 8, 9, 1, 3))
	fmt.Println("La suma de impares con multiargumento es: ", paresImparesVarica("1", 5, 6, 7, 2, 3, 1, 7, 8, 9, 1, 3))
}
