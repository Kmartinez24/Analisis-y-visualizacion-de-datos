package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

/*func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		fmt.Print(rand.Intn(10), " ")
	}
}*/

//Guardar numeros en array

/*const N int = 100

func printArraya(miArray [N]int, n int) {
	fmt.Print("[")
	for i := 0; i < n; i++ {
		fmt.Print(miArray[i], " ")
	}
	fmt.Print("]")
}

func main() {
	var miArray [N]int
	min, max := 10, 50
	n := 20
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		miArray[i] = rand.Intn(max-min+1) + min

	}
	fmt.Println(miArray)

	printArraya(miArray, n)
}*/

/*const N int = 100

func printArraya(miArray [N][N]float64, n, m int) {
	fmt.Println("[")
	for i := 0; i < n; i++ {
		fmt.Print("[")
		for j := 0; j < m; j++ {
			fmt.Print(miArray[i][j], " ")
		}
		fmt.Println("]")
	}
	fmt.Print("]")
}

func main() {
	var miArray [N][N]float64
	n, m := 4, 5
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			numeroRandom := rand.Float64()
			numStr := fmt.Sprintf("%2.2f", numeroRandom)
			miArray[i][j], _ = strconv.ParseFloat(numStr, 64ww)
		}
	}
	//fmt.Println(miArray)
	printArraya(miArray, n, m)
}*/

/*
func printSlice(miSlice []int, n int) {
	fmt.Print("[")
	for _, item := range miSlice {
		fmt.Print(item, " ")
	}
	fmt.Print("]")
}

func main() {
	var miSlice []int
	n, m := 4, 5
	min, max := 10, 50
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			miSlice = append(miSlice, rand.Intn(max-min+1)+min)
		}
	}
	fmt.Println(miSlice)
	printSlice(miSlice, n)
}
*/
func printSlice(miSlice [][]float64) {
	for _, renglon := range miSlice {
		fmt.Println(renglon, " ")
	}

	for _, renglon := range miSlice {
		for _, item := range renglon {
			fmt.Print(item, " ")
		}
		fmt.Println()
	}
}

func main() {
	n, m := 4, 5
	var miSlice = make([][]float64, n)
	for i := 0; i < n; i++ {
		miSlice[i] = make([]float64, m)
	}
	fmt.Println(miSlice)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			numeroRandom := rand.Float64()
			numStr := fmt.Sprintf("%2.2f", numeroRandom)
			miSlice[i][j], _ = strconv.ParseFloat(numStr, 64)
		}
	}
	fmt.Println(miSlice)
	printSlice(miSlice)
}
