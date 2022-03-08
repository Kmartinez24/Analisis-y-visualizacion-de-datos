package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var miMap map[string]int
	miMap = make(map[string]int)
	miMap["key"] = 10
	fmt.Println(miMap)

	elements := make(map[string]string)
	elements["H"] = "Hidrogeno"
	elements["He"] = "Helio"
	elements["Be"] = "Berilio"
	elements["Li"] = "Litio"
	elements["B"] = "Boro"
	elements["C"] = "Carbono"
	elements["N"] = "Nitrogeno"
	elements["O"] = "Oxigeno"
	elements["F"] = "Fluor"
	elements["Ne"] = "Neon"
	fmt.Println(elements)

	/*for clave, valor := range elements {
		fmt.Println(clave, ":", valor)
	}*/
	fmt.Println(elements["Li"])

	/*element, ok := elements["Li"]
	fmt.Println(element, ok)*/

	/*if element, ok := elements["Be"]; ok {
		fmt.Println(element, ok)
	} else {
		fmt.Println("Error todo pete como jota")
	}*/

	/*calif := make(map[string]float64)
	calif["Mario"] = 9.85
	calif["Julia"] = 9.9
	calif["Ana"] = 8.7
	calif["Zapata"] = 6.0
	fmt.Println(calif)
	fmt.Println("Longitud del mapa:", len(calif))*/

	/*numeros := make(map[int]int)
	for i := 0; i < 10; i++ {
		numeros[i] = i + 1
	}
	fmt.Println(numeros)
	numeros[100] = 101
	fmt.Println(numeros)*/
	/*delete(calif, "Zapata")
	fmt.Println(calif)

	nombre, presente := calif["Ana"]
	if presente {
		fmt.Println("La calificacion es:", calif)
	} else {
		fmt.Println("No tiene calificacion")
	}*/
	calif := make(map[int]int)
	var promedio float32
	for i := 0; i < 10; i++ {
		calif[i] = rand.Intn(11-5) + 5
	}
	fmt.Println(calif)
	for _, valor := range calif {
		promedio += float32(valor)
	}
	fmt.Println("El promedio es igual a ", promedio/10)
}
