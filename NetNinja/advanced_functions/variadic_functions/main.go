package main

import "fmt"

func soma(numeros ...int) int {
	total := 0

	for _, numeros := range numeros {
		total += numeros

	}
	return total

}

func escrever(text string, numeros ...int) {
	for _, numeros := range numeros {
		fmt.Println(text, numeros)
	}
}

func main() {
	totalSoma := soma(1, 3, 4, 5, 6)
	fmt.Println(totalSoma)

	escrever("Ola mundo", 1, 2, 3, 4)
}
