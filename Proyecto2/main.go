package main

import (
	"Proyecto2/tools"
	"fmt"
)

func main() {
	num1, num2 := 6, 10
	fmt.Println("Resultado de la suma", tools.Sumar(num1, num2))
	fmt.Println("La resta es: ", tools.Restar(num1, num2))
	fmt.Println("La resta1 es: ", tools.Restar1(num1, num2))
}
