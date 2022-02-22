package Graficos

import "fmt"

func FindMax(nums []int) int {
	var maxnum, numerotemp int
	for _, numero := range nums {
		if numero > numerotemp {
			numerotemp = numero
			maxnum = numerotemp
		}
	}
	return maxnum
}

func GraficaBarras(nums []int, maxnum int, symbol string) {
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
