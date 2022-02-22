package main

import (
	"fmt"
)

/*func potenciaParImpar1(n int) int {
	if n%2 == 0 {
		return n * n
	} else {
		return n * n * n
	}
}*/

/*func potenciaParImpar2(n int) int {
	var res int
	if n%2 == 0 {
		res = n * n
	} else {
		res = n * n * n
	}
	return res
}*/

func potenciaParImpar3(n int) (r int) {
	//x := 10
	if n%2 == 0 {
		r = n * n
	} else {
		r = n * n * n
	}
	return
}

func potencias(n int) (cuad, cubo int) {
	cuad = n * n
	cubo = n * n * n
	return
}

func sumaCuadrados(n int) (suma int) {
	if n >= 5 {
		for i := 0; i < n; i++ {
			suma = suma + (n * n)
		}
	} else {
		n = 0
	}
	return
}

func cuadrado(n int) (cuadrado int) {
	cuadrado = n * n
	return
}

func main() {
	/*//Invocar funcion
	fmt.Println(potenciaParImpar3(2))

	//Asignar funcion
	res := potenciaParImpar3(3)
	fmt.Println(res)*/

	//fmt.Println(potencias(4))

	fmt.Println(sumaCuadrados(5))
	fmt.Println(cuadrado(5))
}

//Programa que obtenga mediante una funcion tipo naked return la suma del cuadrado del 1 hasta n (n enviado como argumento mayor o igual a 5.
//Hacer una funcion naked return para hacer el cuadrado
