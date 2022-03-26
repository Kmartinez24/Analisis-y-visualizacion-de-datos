package main

import "fmt"

/*func convertirMayusculas(palabra *string) {
	*palabra = strings.ToUpper(*palabra)
	fmt.Println("Cadena dentro de la funcion", *palabra)
}

func main() {
	palabra := "hola"
	convertirMayusculas(&palabra)
	fmt.Println("Cadena fuera de la funcion", palabra)

	//closures

	suma := func(x, y int) int { return x + y }
	resta := func(x, y int) int { return x - y }
	sumaResta := func(x, y int) (int, int) { return x + y, x - y }

	fmt.Println("La suma es: ", suma(5, 10))
	fmt.Println("La resta es: ", resta(10, 5))
	r1, r2 := sumaResta(10, 5)
	fmt.Println("La suma es:", r1, "y la resta es", r2)

	x := 10
	incrementar := func() int {
		x++
		return x
	}

	for i := 0; i < 10; i++ {
		fmt.Print(incrementar(), "-")
	}
	fmt.Println()
}*/

//Manejo de estructuras
//registro

type Persona struct {
	nombre      string //campo
	aPaterno    string
	aMaterno    string
	aNacimiento int
}

func despAlumno(al Persona, i byte) {
	fmt.Println(i, "Nombre:", al.nombre)
	fmt.Println(i, "Apellido paterno:", al.aPaterno)
	fmt.Println(i, "Apellido materno:", al.aMaterno)
	fmt.Println(i, "Año de nacimiento:", al.aNacimiento)
}

func main() {
	alumnos := [5]Persona{} //array de tipo persona
	alumnos[0] = Persona{"Juan", "Lopez", "Perez", 2001}
	alumnos[1] = Persona{"Maria", "Perez", "Lopez", 2002}
	alumnos[2] = Persona{"Paco", "Gonzalez", "Jimenez", 2003}
	alumnos[3] = Persona{"Eliza", "Mata", "Lozana", 2000}
	alumnos[4] = Persona{"Juana", "Ruvalcaba", "Rios", 2001}

	for i := 0; i < len(alumnos); i++ {
		fmt.Println(alumnos[i])
	}

	for i := 0; i < len(alumnos); i++ {
		despAlumno(alumnos[i], byte(i+1))
	}

}
