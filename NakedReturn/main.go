package main

import "fmt"

/*
import (
	"fmt"
	"os"
	"strconv"
)

func sumaCuadrados(n int) (suma int) {
	for i := 0; i <= n; i++ {
		if i%2 == 0 {
			suma = suma + (i * i)
		}
	}
	return
}

func sumaCubos(n int) (suma int) {
	for i := 0; i <= n; i++ {
		if i%2 != 0 {
			suma = suma + (i * i * i)
		}
	}
	return
}

func parImpar(n int) int {
	if n%2 != 0 {
		return sumaCubos(n)
	} else {
		return sumaCuadrados(n)
	}

}

func main() {
	Numero, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Hay un error en el codigo")
	} else {
		fmt.Println("El resultado es", parImpar(Numero))
	}
}
*/

func main() {
	for i, j := 1, 10; i <= 10; i, j = i+1, j-1 {
		fmt.Printf("Contador i: %d | Contador j: %d \n", i, j)
	}

	//while
	valid, i, veces := true, 1, 10
	for valid {
		if i == veces {
			valid = false
		}
		fmt.Println(i)
		i++
	}

	i = 1
	for i <= 10{
		fmt.Println(i)
		i++
	}
	 c := 1
	 for{
		 if c > 10{
			 return
		 }
		 fmt.Println(c)
		 c++
	 }
}
