package main

import "fmt"

func diaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	default:
		return "nmr invalido"

	}
}

func diaSemana2(numero int) string {
	var diadaSemana string

	switch {
	case numero == 1:
		diadaSemana  = "Domingo"
	case numero == 2:
		diadaSemana =  "Segunda"
	case numero == 3:
		diadaSemana = "Terça"
	default:
		diadaSemana = "nmr invalido"

	}

	return diadaSemana
}

func main() {
	fmt.Println("Switch")

	dia := diaSemana(1)
	fmt.Println(dia)

	dia2 := diaSemana2(2)
	fmt.Println(dia2)
}
