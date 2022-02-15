package saludo

import "fmt"

//N1 = Primer nombre, N2 = Segundo nombre, N3 = Apellido Paterno, N4 = Apellido Materno, E = Edad, C = Ciudad

func SaludoSimple(N1 string, N2 string){
	fmt.Println("Hola", N1, N2)
}

func SaludoCompleto(N1 string, N2 string, N3 string, N4 string){
	fmt.Println("Hola", N1, N2, N3, N4)
}

func SaludoInvertido(N1 string, N2 string, N3 string, N4 string){
	fmt.Println("Hola", N3, N4, N1, N2)
}

func SaludoEdad(N1 string, E string) {
	fmt.Println("Hola", N1, "de", E, "años")
}

func SaludoCiudad(N1 string, N2 string, N3 string, N4 string, E string, C string){
	fmt.Println("Soy", N1, N2, N3, N4,", vivo en", C, "y tengo", E, "años")
}
