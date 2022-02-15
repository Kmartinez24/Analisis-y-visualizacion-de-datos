package main

import (
	"fmt"
	"os"
)

func main() {
	N := os.Args[1]
	AP := os.Args[2]
	AM := os.Args[3]
	AN := os.Args[4]

	fmt.Println("Hola", AP, AM, N, "naciste el año de ", AN)
}
