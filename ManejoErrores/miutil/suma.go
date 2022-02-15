package miutil

import "fmt"

var numero int = 10
var Numero2 int = 20

func Suma(n1, n2 int) int {
	fmt.Println(numero, Numero2)
	return n1 + n2
}

func sumaCuadrados(n1, n2 int) int {
	return n1*n1 + n2*n2
}
