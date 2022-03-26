package main

/*type Direccion struct {
	calle  string
	numero int
}

type Persona struct {
	nombre    string
	domicilio Direccion
}

func main() {
	var persona1 = new(Persona)
	persona1.nombre = "Ali"
	persona1.domicilio.calle = "bugambilias"
	persona1.domicilio.numero = 333
	fmt.Println("Persona 1:", persona1)

	var persona2 = Persona{"Juan", Direccion{"del cuaco", 234}}
	fmt.Println("Persona 2:", persona2)

	var persona3 = Persona{
		nombre: "Hugo",
		domicilio: Direccion{calle: "Laguna grande",
			numero: 123,
		},
	}
	fmt.Println("Persona 3:", persona3)
}
*/

/*type Usuario struct {
	nombre   string
	apellido string
}

func main() {
	//maps
	mapaString := make(map[string]string)
	mapaString["Nombre"] = "Alex"
	mapaString["Apellido"] = "Mata"

	fmt.Println(mapaString["Nombre"])
	fmt.Println(mapaString["Apellido"])

	//maps con structs
	miMap := make(map[string]Usuario)
	yo := Usuario{"Alex", "Mata"}
	miMap["yo"] = yo
	fmt.Println(miMap["yo"].nombre)
	fmt.Println(miMap["yo"].apellido)
}*/

//Metodos de tipo
/*type num int

func (n1 num) cuadrado() num {
	return n1 * n1
}

func (n1 num) cubo() num {
	return n1 * n1 * n1
}

func main() {
	numero := num(4)
	resCuadrado := numero.cuadrado()
	fmt.Println(resCuadrado)
	resCubo := numero.cubo()
	fmt.Println(resCubo)
}*/

//Struct sencillo con metodo

/*type myStruct struct {
	Nombre string
}

func (n *myStruct) imprimirNombre() string {
	return n.Nombre
}

func main() {
	var miVar1 myStruct
	miVar1.Nombre = "Maria"

	miVar2 := myStruct{Nombre: "Luisa"}

	fmt.Println(miVar2.imprimirNombre(), miVar1.imprimirNombre())
}*/

/*type num int

func (n1 num) suma(n2 num) num {
	return n1 + n2
}

func (n1 num) resta(n2 num) num {
	return n1 - n2
}

func (n1 num) multiplicacion(n2 num) num {
	return n1 * n2
}

func main() {
	numero := num(4)
	suma := numero.suma(5)
	fmt.Println("Suma de", numero, "+ 5 =", suma)
	resta := numero.resta(2)
	fmt.Println("Resta de", numero, "- 2 =", resta)
	multi := numero.multiplicacion(5)
	fmt.Println("Multiplicacion de", numero, "* 5 =", multi)

}*/

//Mediante un truct con dos enteros, hacer los metodos de tipo para que sume, reste y multiplique esos numeros

/*type num int

type numeros struct {
	n1 num
	n2 num
}

func (nu *numeros) suma() num {
	return nu.n1 + nu.n2
}

func (nu *numeros) resta() num {
	return nu.n1 - nu.n2
}

func (nu *numeros) multi() num {
	return nu.n1 * nu.n2
}

func main() {
	var miVar numeros
	miVar.n1 = 10
	miVar.n2 = 5
	suma := miVar.suma()
	resta := miVar.resta()
	multipicacion := miVar.multi()
	fmt.Println("Resultado suma", suma)
	fmt.Println("Resultado resta", resta)
	fmt.Println("Resultado multiplicacion", multipicacion)
}*/

/*type entero int
type decimal float64

func (n1 entero) suma(n2 entero) entero {
	return n1 + n2
}

func (n1 decimal) suma(n2 decimal) decimal {
	return n1 + n2
}

func main() {
	var numero1 entero
	numero1 = 5
	var numero2 decimal
	numero2 = 4.5
	fmt.Println(numero1.suma(6))
	fmt.Println(numero2.suma(5.2))

}*/
