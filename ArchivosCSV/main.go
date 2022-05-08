package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func obtenerMes(n int) (mes string) {
	switch n {
	case 1:
		mes = "Enero"
		break
	case 2:
		mes = "Febrero"
		break
	case 3:
		mes = "Marzo"
		break
	case 4:
		mes = "Abril"
		break
	case 5:
		mes = "Mayo"
		break
	case 6:
		mes = "Junio"
		break
	case 7:
		mes = "Julio"
		break
	case 8:
		mes = "Agusto"
		break
	case 9:
		mes = "Septiembre"
		break
	case 10:
		mes = "Octubre"
		break
	case 11:
		mes = "Noviembre"
		break
	case 12:
		mes = "Diciembre"
		break
	default:
		mes = ""
	}
	return
}

func main() {
	numTemp := 4

	numMesesStr := os.Args[1]
	minStr := os.Args[2]
	maxStr := os.Args[3]

	numMeses, _ := strconv.Atoi(numMesesStr)
	min, _ := strconv.Atoi(minStr)
	max, _ := strconv.Atoi(maxStr)

	var Mes, Sem1, Sem2, Sem3, Sem4 []string
	var sliceFinal [][]string

	rand.Seed(time.Now().UnixNano())
	renglonMes := ""
	for i := 1; i <= numMeses; i++ {
		fmt.Printf("%s\t", obtenerMes(i))
		Mes = append(Mes, fmt.Sprintf(obtenerMes(i)))
	}
	fmt.Println()
	fmt.Println(Mes)
	fmt.Println(renglonMes)

	for i := 1; i <= numTemp; i++ {

		for j := 0; j < numMeses; j++ {

			tempRandom := rand.Intn(max-min) + min
			if i == 1 {
				Sem1 = append(Sem1, fmt.Sprintf("%d", tempRandom))
			}
			if i == 2 {
				Sem2 = append(Sem2, fmt.Sprintf("%d", tempRandom))
			}
			if i == 3 {
				Sem3 = append(Sem3, fmt.Sprintf("%d", tempRandom))
			}
			if i == 4 {
				Sem4 = append(Sem4, fmt.Sprintf("%d", tempRandom))

			}
		}
	}

	sliceFinal = append(sliceFinal, Mes)
	sliceFinal = append(sliceFinal, Sem1)
	sliceFinal = append(sliceFinal, Sem2)
	sliceFinal = append(sliceFinal, Sem3)
	sliceFinal = append(sliceFinal, Sem4)
	fmt.Println(sliceFinal)

	create, err := os.Create("./data.csv")
	if err != nil {
		fmt.Println(err)
	}

	writer := csv.NewWriter(create)

	err = writer.WriteAll(sliceFinal)
	if err != nil {
		fmt.Println(err)
	}
}
