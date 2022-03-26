package main

/*type Persona struct {
	Nombre string
	Edad   int
}

func (P *Persona) setNombre(nombre string) {
	(*P).Nombre = nombre
}

func (P Persona) mostrar() {
	fmt.Println("Nombre:", P.Nombre)
	fmt.Println("Edad:", P.Edad)
}

func main() {
	estudiante := Persona{"Alex", 45}
	fmt.Println(estudiante)
	e := &estudiante
	e.setNombre("Alexander")
	fmt.Println(estudiante)
	estudiante.mostrar()
}
*/
/*
type Detalles struct {
	Nombre  string
	Edad    int
	Genero  string
	Salario int
}

type Empleado struct {
	id int
	Detalles
}

func (d Detalles) calcularSalario(dias int) int {
	return d.Salario * dias
}

func (e Empleado) mostrar() {
	fmt.Println("ID Empleado:", e.id)
	fmt.Println("Nombre Empleado:", e.Nombre)
	fmt.Println("Edad: Empleado:", e.Edad)
	fmt.Println("Genero Empleado:", e.Genero)
	fmt.Println("Salario Empleado:", e.Salario)
}

func main() {
	a := Empleado{
		id: 12345,
		Detalles: Detalles{
			Nombre:  "Alex",
			Edad:    34,
			Genero:  "No binario",
			Salario: 950,
		},
	}
	a.mostrar()
	fmt.Println("Sueldo total: ", a.calcularSalario(15), "pesos")

}*/

//Funciones como campos de una estructura

/*type operacion func(int, int) int

type Calculadora struct {
	num1 int
	num2 int
	op   func(int, int) int
}

func main() {
	res := Calculadora{
		num1: 5,
		num2: 10,
		op: func(n1 int, n2 int) int {
			return n1 + n2
		},
	}

	resta := Calculadora{
		num1: 5,
		num2: 10,
		op: func(n1 int, n2 int) int {
			return n1 - n2
		},
	}

	multi := Calculadora{
		num1: 5,
		num2: 10,
		op: func(n1 int, n2 int) int {
			return n1 * n2
		},
	}

	fmt.Println(res.num1, "+", res.num2, "=", res.op(res.num1, res.num2))
	fmt.Println(resta.num1, "-", resta.num2, "=", resta.op(resta.num1, resta.num2))
	fmt.Println(multi.num1, "*", multi.num2, "=", multi.op(multi.num1, multi.num2))
}*

//Interfaces
*/

/* func imprimir(i interface{}) {
	fmt.Println(i)
}

func main() {
	imprimir(1)
	imprimir(1.5)
	imprimir("Hola")
	imprimir([]int{1, 2, 3, 4, 5})
}*/

/*func add(a, b interface{}) (interface{}, error) {
	typeA := reflect.TypeOf(a).Kind()
	typeB := reflect.TypeOf(b).Kind()
	if typeB != typeA {
		return nil, errors.New("ambos parametros deben ser del mismo tipo")
	}
	switch typeA {
	case reflect.Int:
		return a.(int) + b.(int), nil
	case reflect.Int8:
		return a.(int8) + b.(int8), nil
	case reflect.Float32:
		return a.(float32) + b.(float32), nil
	default:
		return nil, errors.New("parametros invalidos")
	}
}

func main() {
	var n1, n2 int8 = 5, 6
	res, err := add(4, 5.1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("La suma es:", res)
	}
	fmt.Println(add(n1, n2))
	fmt.Println(add(3, 6))
}*/

//Reflections and Generics
/*func add[T int | int8 | int16 | int32 | int64 | float32 | float64 | string](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(add(5.1, 6.1))
	fmt.Println(add(5, 6))
	fmt.Println("Hola", "Mundo")
	fmt.Println(add(int8(5), int8(2)))
}*/
