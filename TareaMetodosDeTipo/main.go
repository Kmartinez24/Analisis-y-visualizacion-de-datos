package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type entero int

//1. Se desea obtener la suma de todos los números de una lista de números variables enviados al método. La llamada
//puede ser por ejemplo: serie.sum(1,2,3,4,5,6,7,8,9,10) cuyo resultado es 55.
func (n1 entero) suma(nums []int) (resultado entero) {
	resultado = 0
	for _, num := range nums {
		tempnum := entero(num)
		resultado += tempnum
	}
	return
}

//2. Obtener la suma de todos los números pares e impares de manera independiente en un solo método.
func (n1 entero) paresImpares(nums []int) (resultadoPares, resultadoImpares entero) {
	resultadoPares, resultadoImpares = 0, 0
	for _, num := range nums {
		tempnum := entero(num)
		if tempnum%2 == 0 {
			resultadoPares += tempnum
		} else {
			resultadoImpares += tempnum
		}
	}
	return
}

//3. Obtener la suma de los cuadrados de todos los números
func (n1 entero) sumaCuadrados(nums []int) (resultado entero) {
	resultado = 0
	for _, num := range nums {
		tempnum := entero(num)
		resultado += tempnum * tempnum
	}
	return
}

//4. Obtener el promedio de todos los números.
func (n1 entero) promedio(nums []int) (promedio float64) {
	var resultado entero = 0
	for _, num := range nums {
		tempnum := entero(num)
		resultado += tempnum
	}
	resultadoF := float64(resultado)
	longitudF := float64(len(nums))
	promedio = resultadoF / (longitudF)
	promedio = math.Round(promedio*100) / 100
	return
}

//5. Obtener la suma de los cubos de todos los números.
func (n1 entero) sumaCubos(nums []int) (resultado entero) {
	resultado = 0
	for _, num := range nums {
		tempnum := entero(num)
		resultado += tempnum * tempnum * tempnum
	}
	return
}

func main() {
	var numero1 entero
	var nums []int
	rand.Seed(time.Now().UnixNano())
	long, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("El argumento no es un numero")
	} else {
		for i := 0; i < long; i++ {
			nums = append(nums, rand.Intn(10)+1)
		}
		fmt.Println("Los", long, "numeros generados son", nums)
		fmt.Println("Resultado de la suma:", numero1.suma(nums))
		resPares, resImpares := numero1.paresImpares(nums)
		fmt.Println("Resultado de suma de pares:", resPares)
		fmt.Println("Resultado de suma de impares:", resImpares)
		fmt.Println("Resultado de suma de cuadrados", numero1.sumaCuadrados(nums))
		fmt.Println("Promedio de todos los numeros", numero1.promedio(nums))
		fmt.Println("Resultado de suma de todos los cubos", numero1.sumaCubos(nums))
	}
}
