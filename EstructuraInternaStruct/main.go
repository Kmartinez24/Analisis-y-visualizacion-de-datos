package main

import (
	"fmt"
	"reflect"
)

/*type Dir struct {
	Colonia string
	Calle   string
	Numero  int
}

type Persona struct {
	Nombre      string
	ANacimiento int
	Domicilio   Dir
}

func main() {
	P := Persona{"Juan Perez", 1995,
		Dir{"Bonanza", "La Predegoza", 123}}

	r := reflect.ValueOf(P)
	fmt.Println("R=", r)
	fmt.Println("Valor del string=", r.String())
	fmt.Println("-----------------------")
	iType := r.Type()
	fmt.Println("i Type:", iType)

	fmt.Println("-----------------------")
	for i := 0; i < r.NumField(); i++ {
		fmt.Printf("Campo:\t[%s]", iType.Field(i).Name)
		fmt.Printf("\tTipo [%s]", r.Field(i).Type())
		fmt.Printf("\tValor [%v]\n", r.Field(i).Interface())
		fmt.Println()

	}
	fmt.Println("-----------------------")
	fmt.Println(r.Field(2).Field(0))
	fmt.Println("-----------------------")
	for i := 0; i < r.NumField(); i++ {
		k := reflect.TypeOf(r.Field(i).Interface()).Kind()
		if k == reflect.Struct {
			jType := r.Field(i).Type()
			for j := 0; j < r.Field(i).NumField(); j++ {
				fmt.Printf("Campo:\t[%s]", jType.Field(j).Name)
				fmt.Printf("\tTipo [%s]", r.Field(i).Field(j).Type())
				fmt.Printf("\tValor [%v]\n", r.Field(i).Field(j).Interface())
				fmt.Println()
			}
		}
	}
}
*/

type Numeros struct {
	Entero   int
	Cadena   string
	Flotante float64
}

func main() {
	N := Numeros{1, "Temperatura", 3.1416}
	fmt.Println("N:", N)

	r := reflect.ValueOf(&N).Elem()
	fmt.Println("Valor string:", r.String())

	typeofN := r.Type()
	fmt.Println(typeofN)
	for i := 0; i < r.NumField(); i++ {
		f := r.Field(i)
		tofN := typeofN.Field(i).Name
		fmt.Printf("%d: %s %s = %v\n", i, tofN, f.Type(), f.Interface())
		k := reflect.TypeOf(r.Field(i).Interface()).Kind()
		if k == reflect.Int {
			r.Field(i).SetInt(-100)
		} else if k == reflect.String {
			r.Field(i).SetString("Humedad")
		}
	}
	fmt.Println("*******************")
	fmt.Println("N", N)

	N.Cadena = "Hola mundo"
	fmt.Println("N:", N)
}
