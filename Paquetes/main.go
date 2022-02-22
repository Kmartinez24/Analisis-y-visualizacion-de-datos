package main

import (
	"Paquetes/calcfloats"
	"Paquetes/calcints"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	//Operaciones con float32
	n1 := rand.Float32()
	n2 := rand.Float32()
	fmt.Println("Los numeros a utilizar seran", n1, n2)
	fmt.Println("Suma de float", calcfloats.Suma(n1, n2))
	fmt.Println("Resta de float", calcfloats.Resta(n1, n2))
	fmt.Println("Resta alrevez de float", calcfloats.Resta2(n1, n2))
	fmt.Println("Multiplicacion de float", calcfloats.Multiplicacion(n1, n2))
	fmt.Println("Division de float", calcfloats.Division(n1, n2))

	//Operaciones con int
	n3 := rand.Intn(10-1) + 1
	n4 := rand.Intn(10-1) + 1
	fmt.Println("\nLos numeros a utilizar seran", n3, n4)
	fmt.Println("Suma de ints", calcints.Suma(n3, n4))
	fmt.Println("Resta de ints", calcints.Resta(n3, n4))
	fmt.Println("Resta al revez de ints", calcints.Resta2(n3, n4))
	fmt.Println("Multiplicacion de ints", calcints.Suma(n3, n4))
	fmt.Println("Division de ints", calcints.Division(n3, n4))

}
