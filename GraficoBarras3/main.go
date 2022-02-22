package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	//Grafico de barras que recibe los argumentos desde la terminal
	var nums []int
	var dato int
	var symbol string

	for i := 0; i < len(os.Args); i++ {
		if i > 1 {
			dato, _ = strconv.Atoi(os.Args[i])
			nums = append(nums, dato)
			fmt.Println(dato)
		}
	}
	symbol = os.Args[1]
	fmt.Println(nums)
	fmt.Println("Numero maximo = ", findMax(nums))
	graficaBarras(nums, findMax(nums), symbol)

}

func findMax(nums []int) int {
	var maxnum, numerotemp int
	for _, numero := range nums {
		if numero > numerotemp {
			numerotemp = numero
			maxnum = numerotemp
		}
	}
	return maxnum
}

func graficaBarras(nums []int, maxnum int, symbol string) {
	for j := maxnum; j > 0; j-- {
		fmt.Printf("%d", j)
		fmt.Printf("|")
		for i := 0; i < len(nums); i++ {
			if nums[i] >= j {
				fmt.Printf("%s", symbol)
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
	for f := 0; f < len(nums)+2; f++ {
		fmt.Printf("-")
	}
}
