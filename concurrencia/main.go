package main

import (
	"fmt"
	"sync"
	"time"
)

func cuadrado(num int) {
	time.Sleep(time.Second * 3)
	fmt.Println(num * num)
	wg.Done()
}

func cubo(num int) {
	time.Sleep(time.Second * 3)
	fmt.Println(num * num * num)
	wg.Done()
}

var wg = sync.WaitGroup{}

func main() {
	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	wg.Add(len(numeros))
	fmt.Println("Realizando calculos")
	inicio := time.Now()
	for _, numero := range numeros {
		go cuadrado(numero)
	}
	wg.Wait()
	tEjecucion := time.Since(inicio)
	fmt.Println("Tiempo total de ejecucion:", tEjecucion)

}
