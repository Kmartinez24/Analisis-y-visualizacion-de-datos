package main

import (
	"fmt"
)

/*
import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

		PRIMER CODIGO
func cuadradoPar2(n int) (int, bool) {
	if n%2 == 0 {
		return n * n, true
	} else {
		return 0, false
	}
}

func cuadradoPar(n int) (int, error) {
	if n%2 == 0 {
		return n * n, nil
	} else {
		return 0, errors.New("tiene que ser par")
	}
}

func main() {
	numero, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("El error es: ", err)
	} else {
		cuad, err := cuadradoPar(numero)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("El cuadrado es: ", cuad)
		}
	}



}*/
/*
func main() {
	fmt.Println(miutil.Numero2)
	miutil.Suma(1, 2)
	fmt.Println(miutil.Resta(3, 4))

}
*/

/*var n1 int = 30
var n2 rune

func mostrarDatos() {
	fmt.Println("desde mostrarDatos:", n1, n2, "Tipo:", reflect.TypeOf(n2))
}

func main() {
	var n1 int = 10
	var n2 rune
	n3 := 20

	fmt.Println(n1, n2, n3)
	mostrarDatos()
}
*/
/*
func main() {
	var n1, n2 = 10, 20
	if n1 < n2 {
		n3 := 5
		fmt.Println(n3, "Es menor")
	} else {
		fmt.Println(n3)
		fmt.Println("Es mayor")
	}
	fmt.Println(n3)
}
*/

func cuadradoPar(n int) (int, bool) {
	if n%2 == 0 {
		return n * n, true
	} else {
		return 0, false
	}
}

func main() {
	res1, ok1 := cuadradoPar(3)
	if ok1 {
		fmt.Println("El cuadrado es: ", res1)
	} else {
		fmt.Println("Error...")
	}
	fmt.Println(res1, ok1)
	if res2, ok2 := cuadradoPar(3); ok2 {
		fmt.Println("El cuadrado es: ", res2)
		fmt.Println(res2, ok2)
	} else {
		fmt.Println("Error...")
		fmt.Println(res2, ok2)
	}
}
