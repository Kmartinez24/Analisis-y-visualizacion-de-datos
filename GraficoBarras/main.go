package main

import (
	"fmt"
)

func main() {
	var nums []int
	var cdatos, n int
	var symbol string

	fmt.Printf("Simbolo para rellenar: ")
	fmt.Scanln(&symbol)
	fmt.Printf("Cuantos datos desea ingresar? ")
	fmt.Scanln(&cdatos)
	for i := 0; i < cdatos; i++ {
		fmt.Printf("Numero de la posicion %d: ", i+1)
		fmt.Scan(&n)
		nums = append(nums, n)
	}
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
			//fmt.Printf("%d", j)
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
