package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops")
	//i := 0

	//for i =< 10 {
	//	i++
	// fmt.Println("Incrementando i")
	// 	time.Sleep(time.Second)
	// }

	//fmt.Println(i)

	// for j := 0; j < 10; j++ {
	// 	fmt.Println("Incrementando j", j)
	// 	time.Sleep(time.Second)

	// }

	nomes := [3]string{"joao", "david", "pedro"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	user := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Dias",
	}


	for _, valor := range user {
		fmt.Println(valor)
	}
}
