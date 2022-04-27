package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

/*import "fmt"

var count = 0

func conteoRecursivo(n int) {
	count++
	if count <= 10 {
		fmt.Println(count)
		conteoRecursivo(n)
	}
}

func main() {
	conteoRecursivo(15)
}*/

/*func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	fmt.Println("El factorial de 5 es:", factorial(5))
}
*/

//Archivos CSV
/*
Enero, Febrero, Marzo <-encabezado
15.25, 13.25, 45.25 <- datos*/

/*func main() {
	create, err := os.Create("./data.csv")
	if err != nil {
		fmt.Println(err)
	}
	writer := csv.NewWriter(create)
	var data = [][]string{
		{"Enero", "Febrero", "Marzo"},
		{"15.25", "13.25", "14.25"},
		{"14.25", "11.25", "16.25"},
		{"15.25", "12.25", "18.25"},
		{"16.25", "16.25", "19.25"},
		{"17.25", "17.25", "20.25"},
	}
	err = writer.WriteAll(data)
	if err != nil {
		fmt.Println(err)
	}
}*/

/*func main() {
	file, err := os.Open("./data.csv")
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(file)
	all, err := reader.ReadAll()
	if err != nil {
		return
	}
	fmt.Println(all)
	for _, all := range all {
		fmt.Println(all)
	}
	for _, strings := range all {
		for _, s := range strings {
			fmt.Println(s)
		}
	}
}*/

func main() {
	tempBaseStr := os.Args[1]
	numTempStr := os.Args[2]
	numMesesStr := os.Args[3]
	tbase, _ := strconv.ParseFloat(tempBaseStr, 64)
	numTemp, _ := strconv.Atoi(numTempStr)
	numMeses, _ := strconv.Atoi(numMesesStr)

	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= numMeses; i++ {
		fmt.Printf("%d\t", i)
	}
	fmt.Println()
	renglon := ""
	for i := 1; i <= numTemp; i++ {
		for j := 0; j < numMeses; j++ {
			tempRandom := tbase + rand.Float64()
			renglon = renglon + fmt.Sprintf("%2.2f\t", tempRandom)
		}
		fmt.Println(renglon)
		renglon = ""
	}
}
