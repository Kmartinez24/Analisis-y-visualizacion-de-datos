package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	calif := make(map[int]int)
	var promedio float32
	for i := 0; i < 10; i++ {
		calif[i+1] = rand.Intn(11-5) + 5
	}
	fmt.Println(calif)
	for _, valor := range calif {
		promedio += float32(valor)
	}
	fmt.Println("El promedio es igual a ", promedio/10)
}
