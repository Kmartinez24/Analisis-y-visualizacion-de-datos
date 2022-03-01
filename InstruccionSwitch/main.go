package main

func main() {
	//fruta1 := "fresa"
	//fruta2 := "manzana"

	//seleccion := os.Args[1]

	/*switch seleccion {
	case fruta1:
		fmt.Println("Seleccionaste", seleccion)
	case fruta2:
		fmt.Println("Seleccionaste", seleccion)
	default:
		fmt.Println(seleccion, "no esta disponible")
	}*/

	/*switch true {
	case fruta1 == seleccion:
		fmt.Println("Seleccionaste", seleccion)
	case fruta2 == seleccion:
		fmt.Println("Seleccionaste", seleccion)
	case 5 > 2:
		fmt.Println("5 es mayor que 2")
	default:
		fmt.Println(seleccion, "no esta disponible")
	}*/

	/*letra := []int32{'a', '0', '+', 'e', '9', '*', 'u', 'b'}
	tam := len(letra)
	for i := 0; i < tam; i++ {
		switch letra[i] {
		case 'a', 'e', 'i', 'o', 'u':
			fmt.Printf("%c es una vocal\n", letra[i])
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			fmt.Printf("%c es un número\n", letra[i])
		default:
			fmt.Println("Es otro símbolo")
		}
	}*/
	/*letra := []byte{'a', '0', '+', 'e', '9', '*', 'u', 'b'}
	tam := len(letra)
	for i := 0; i < tam; i++ {
		l := letra[i]
		switch true {
		case l == 'a' || l == 'e' || l == 'i' || l == 'o' || l == 'u':
			fmt.Printf("%c es una vocal\n", letra[i])
		case l >= 'b' && l <= 'z':
			fmt.Printf("%c es una consonante\n", letra[i])
		case l >= '0' && l <= '9':
			fmt.Printf("%c es un número\n", letra[i])
		default:
			fmt.Println("Es otro símbolo")
		}
	}*/
}
