package main

import "os"

func main() {
	Nombre := []string{"K", "i", "y", "o", "s", "h", "i"}
	Nombre2 := []string{"A", "l", "i"}
	ApPa := []string{"M", "a", "r", "t", "i", "n", "e", "z"}
	ApMa := []string{"R", "e", "n", "t", "e", "r", "i", "a"}
	for i := 0; i < 4; i++{
		switch i {
		case 0:
			tam := len(Nombre)
			for j := 0; j < tam; j++{
				print(Nombre[j])
			}
			print("\n")
		case 1:
			tam := len(Nombre2)
			for j := 0; j < tam; j++{
				print(Nombre2[j])
			}
			print("\n")
		case 2:
			tam := len(ApPa)
			for j := 0; j < tam; j++{
				print(ApPa[j])
			}
			print("\n")
		case 3:
			tam := len(ApMa)
			for j := 0; j < tam; j++{
				print(ApMa[j])
			}
			print("\n")
		default:
			println("Error inesperado")
			os.Exit(0)
		}
	}
}

