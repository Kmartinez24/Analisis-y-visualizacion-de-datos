package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sumar(a, b, c int) (res int) {
	res = a + b + c
	return
}

//variadicas
func suma2(nums ...int) (res int) {
	for _, num := range nums {
		res += num
	}
	return
}

func sumarpares(pares []int) (sum int) {
	for _, num := range pares {
		sum += num
	}
	return
}

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
	/*fmt.Println("Suma de pares es :", sumarpares(numerosPares))
	fmt.Println("La suma es:", sumar(1, 2, 3))
	fmt.Println("La suma variadica es:", suma2(1, 2, 3))*/
	//Funcion que se llame pares impares con argumento slice y bandera
	//Lo mismo pero con varios argumentos y una bandera
	fmt.Println("La suma de numeros con slice es: ", paresImparesSlice(numeros, "1"))
	fmt.Println("La suma de numeros con multiargumento es: ", paresImparesVarica("1", 5, 6, 7, 2, 3, 1, 7, 8, 9, 1, 3))
}
