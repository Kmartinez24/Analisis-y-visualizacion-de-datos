package main

import (
	"fmt"
	"os"
	"strconv"
)

func calcularEdad(n int) (int, bool) {
	if n < 2022 {
		return 2022 - n, true
	} else {
		return 0, false
	}
}

func main() {
	N := os.Args[1]
	AP := os.Args[2]
	AM := os.Args[3]
	AN := os.Args[4]

	ANint, _ := strconv.Atoi(AN)       //ASCII to integer sin necesidad de error
	Edad, check := calcularEdad(ANint) //Mando a llamar calcular edad que al mismo tiempo revisa que la edad sea menor que 2022
	if check {
		fmt.Println("Hola", AP, AM, N, ", tienes", Edad, "años de edad")
	} else {
		fmt.Println("La edad es mayor o igual al año actual")
	}
}
