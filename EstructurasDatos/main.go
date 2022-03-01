package main

import "fmt"

func main() {
	/*var array [2]int
	var slice []int
	fmt.Println(array)
	array[0] = 5
	array[1] = 10
	fmt.Println(reflect.TypeOf(array))
	fmt.Printf("%T", slice)*/

	/*var miArray [10]int
	for i := 0; i < 10; i++ {
		miArray[i] = i + 1
	}
	fmt.Println(miArray)

	for i, item := range miArray {
		fmt.Println(i, ":", item)
	}*/
	/*diasSemana := []string{"lunes",
		"martes",
		"miercoles",
		"jueves",
		"viernes",
		"sabado",
		"domingo",
	}
	fmt.Println("Cantidad de dias: ", len(diasSemana))
	for _, dia := range diasSemana {
		fmt.Println(dia, "|")
	}*/
	/*miVector := [12]int{3: 6, 6: 12, 9: 18}
	fmt.Println(miVector)*/
	/*var vacio []int
	slice := make([]int, 10)
	for i := 1; i <= 10; i++ {
		vacio = append(vacio, i)
	}

	for i, _ := range slice {
		slice[i] = i + 1
	}
	fmt.Println(vacio)
	fmt.Println(slice)*/
	slice1 := []int{1, 2, 3, 4, 5, 6}
	slice2 := make([]int, 5)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
}
