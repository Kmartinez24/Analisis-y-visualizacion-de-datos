package main

import "fmt"

func main() {
	america := []string{"Mexico", "Brasil", "Chile", "Peru", "Bolivia"}
	europa := []string{"España", "Portugal", "Dinamarca", "Francia", "Ucrania"}
	intercalados := make([]string, 10)
	pos := 0
	for i := 0; i < 5; i++ {
		intercalados[pos] = europa[i]
		pos += 1
		intercalados[pos] = america[i]
		pos += 1
	}
	fmt.Println(intercalados)

	var intercaladosF []string
	for i := 0; i < 5; i++ {
		intercaladosF = append(intercaladosF, europa[i])
		intercaladosF = append(intercaladosF, america[i])
	}
	fmt.Println(intercaladosF)
}
