package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type pago struct {
	mes     int
	capital float64
	interes float64
	total   float64
	saldo   float64
}

func obtenerPorcentaje(cantidad, interes float64) (porcentaje float64) {
	porcentaje = cantidad * (interes / 100)
	return
}

func calcularInteres(producto string, plazo int) (interes float64) {
	if producto == "efectivo" || producto == "muebles" || producto == "inmuebles" {
		if producto == "efectivo" {
			if plazo >= 3 && plazo <= 12 {
				interes = 2
			} else {
				fmt.Println("Plazo invalido")
				os.Exit(4)
			}
		}
		if producto == "muebles" {
			if plazo >= 6 && plazo <= 24 {
				interes = 2.5
			} else {
				fmt.Println("Plazo invalido")
				os.Exit(4)
			}
		}
		if producto == "inmuebles" {
			if plazo >= 12 && plazo <= 60 {
				interes = float64(interesInmuebles(plazo))
			} else {
				fmt.Println("Plazo invalido")
				os.Exit(4)
			}
		}
	} else {
		fmt.Println("Producto invalido")
		os.Exit(4)
	}
	return
}

func interesInmuebles(plazo int) (interes int) {
	if plazo <= 0 {
		fmt.Println("Error en plazo")
		os.Exit(4)
	}
	if plazo > 0 && plazo <= 12 {
		interes = 5
	}
	if plazo > 12 && plazo <= 24 {
		interes = 4
	}
	if plazo > 24 && plazo <= 36 {
		interes = 3
	}
	if plazo > 36 && plazo <= 48 {
		interes = 2
	}
	if plazo > 48 && plazo <= 60 {
		interes = 1
	}
	return
}

func interesFijo(plazo int, cantidad float64, interes float64) {
	var iTotal float64

	pTotal := cantidad + obtenerPorcentaje(cantidad, interes)
	pTotal = math.Round(pTotal*100) / 100

	pMes := cantidad / float64(plazo)
	pMes = math.Round(pMes*100) / 100
	iTotal = obtenerPorcentaje(pMes, interes) * float64(plazo)

	iMes := iTotal / float64(plazo)
	iMes = math.Round(iMes*100) / 100

	pagos := make([]pago, plazo)
	saldo := (pMes + iMes) * float64(plazo)
	pTotal = saldo
	for i := range pagos {
		saldo -= pMes + iMes
		saldo = math.Round(saldo*100) / 100
		pagos[i] = pago{
			mes:     i + 1,
			capital: pMes,
			interes: iMes,
			total:   math.Round((pMes+iMes)*100) / 100,
			saldo:   saldo,
		}
	}
	for i := range pagos {
		fmt.Println("Mes", pagos[i].mes, "Capital", pagos[i].capital, "Interes", pagos[i].interes, "Total", pagos[i].total, "Saldo", pagos[i].saldo)
	}
	fmt.Println("Con un interes porcentual mensual de", interes, "% que equivale a $", iTotal, ", dando como resultado una deuda de $", pTotal)
}

func interesDecreciente(plazo int, cantidad float64, interes float64) {
	var iTotal, pTotal float64
	pMes := cantidad / float64(plazo)
	pMes = math.Round(pMes*100) / 100

	iMes := obtenerPorcentaje(pMes, interes)
	iMes = math.Round(iMes*100) / 100

	pagos := make([]pago, plazo)
	saldo := pMes * float64(plazo)
	for i := range pagos {
		saldo -= pMes
		saldo = math.Round(saldo*100) / 100
		pagos[i] = pago{
			mes:     i + 1,
			capital: pMes,
			interes: iMes,
			total:   math.Round((pMes+iMes)*100) / 100,
			saldo:   saldo,
		}
		iTotal += iMes
		pTotal += pMes + iMes
		iMes = obtenerPorcentaje(saldo/(float64(plazo)-float64(i)), interes)
		iMes = math.Round(iMes*100) / 100
	}
	iTotal = math.Round(iTotal*100) / 100
	pTotal = math.Round(pTotal*100) / 100

	for i := range pagos {
		fmt.Println("Mes", pagos[i].mes, "Capital", pagos[i].capital, "Interes", pagos[i].interes, "Total", pagos[i].total, "Saldo", pagos[i].saldo)
	}
	fmt.Println("Con un interes inicial mensual de", interes, "% que equivale a $", iTotal, ", dando como resultado una deuda de $", pTotal)
}

func main() {
	fmt.Println("Producto                   | Plazo mínimo | Plazo máximo | Interés")
	fmt.Println("----------------------------------------------------------------------------------------------------------------------------------------")
	fmt.Println("Efectivo                   |  3 meses     |   1 año      | 2% mensual")
	fmt.Println("Bienes muebles             |  6 meses     |   2 años     | 2.5% mensual")
	fmt.Println("Bienes inmuebles           | 12 meses     |   5 años     | 5% mensual por un año disminuyendo")
	fmt.Println("                           |              |              |  en 1% cada año, es decir que si el")
	fmt.Println("                           |              |              |  crédito es a 2 años, el interés será del 4%")

	tempPlazo, errConv := strconv.Atoi(os.Args[3])
	cantidad, errConv3 := strconv.ParseFloat(os.Args[2], 64)
	tInteres := strings.ToLower(os.Args[4])
	producto := strings.ToLower(os.Args[1])

	if errConv != nil || errConv3 != nil {
		fmt.Println("Error en la conversion de datos, por favor revise la entrada")
		os.Exit(5)
	}
	interes := calcularInteres(producto, tempPlazo)

	if tInteres == "fijo" {
		interesFijo(tempPlazo, cantidad, interes)
	}
	if tInteres == "decreciente" {
		interesDecreciente(tempPlazo, cantidad, interes)
	}
	if tInteres != "decreciente" && tInteres != "fijo" {
		fmt.Println("Tipo de interes invalido")
		os.Exit(7)
	}

}
