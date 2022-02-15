package main

import (
	"CursoGo2021/saludo"
	"os"
	"fmt"
)

func main() {
	Nombre1 := os.Args[1]
	Nombre2 := os.Args[2]
	ApPa := os.Args[3]
	ApMa := os.Args[4]
	Edad := os.Args[5]
	Ciudad := os.Args[6]
	fmt.Println("1. Despliega el mensaje \"Hola <tu_nombre>\", reemplazando tu_nombre por tu nombre completo sin apellidos.")
	saludo.SaludoSimple(Nombre1, Nombre2)
	fmt.Println("") //Print vacio para salto de linea

	fmt.Println("2. Despliega tu nombre completo empezando por Nombre(s), Apellido Paterno y Apellido Materno.")
	saludo.SaludoCompleto(Nombre1, Nombre2, ApPa, ApMa)
	fmt.Println("") //Print vacio para salto de linea

	fmt.Println("3. Ahora despliega tu nombre completo, empezando por Apellido Paterno, Apellido Materno y Nombre.")
	saludo.SaludoInvertido(Nombre1, Nombre2, ApPa, ApMa)
	fmt.Println("") //Print vacio para salto de linea

	fmt.Println("4. Despliega solo tu primer nombre y edad.")
	saludo.SaludoEdad(Nombre1, Edad)
	fmt.Println() //Print vacio para salto de linea

	fmt.Println("5. Despliega el mensaje, \"Hola, soy <tu_nombre>, vivo en (tu_ciudad) y tengo (tu_edad) años.")
	saludo.SaludoCiudad(Nombre1, Nombre2, ApPa, ApMa, Edad, Ciudad)
	fmt.Println() //Print vacio para salto de linea
}