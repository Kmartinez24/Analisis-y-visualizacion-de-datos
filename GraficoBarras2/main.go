package main

import (
	"GraficoBarras/Graficos"
	"GraficoBarras/calcfloats"
	"GraficoBarras/calcints"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func muestraOperaciones() {
	rand.Seed(time.Now().UnixNano())
	//Operaciones con float32
	n1 := rand.Float32()
	n2 := rand.Float32()
	fmt.Println("\nLos numeros a utilizar seran", n1, n2)
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

func obtenerSlice() (nums []int, symbol string) {
	var dato int
	for i := 0; i < len(os.Args); i++ {
		if i > 1 {
			dato, _ = strconv.Atoi(os.Args[i])
			nums = append(nums, dato)
		}
	}
	symbol = os.Args[1]
	return
}

func main() {
	/*Grafico de barras que recibe los argumentos desde la terminal, el primer valor enviado desde la terminal debe ser
	el simbolo con el que deseas graficar, todo lo siguiente son los datos a graficar*/
	nums, symbol := obtenerSlice()
	fmt.Println(nums)
	fmt.Println("Numero maximo = ", Graficos.FindMax(nums))
	Graficos.GraficaBarras(nums, Graficos.FindMax(nums), symbol)

	//Invocamiento de funcion para imprimir el uso de operaciones de int y float
	muestraOperaciones()
}
