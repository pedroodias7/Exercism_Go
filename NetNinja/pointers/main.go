package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var variavel2 = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++

	fmt.Println(variavel1, variavel2)

	//pointers memory reference
	var variavel3 int
	var pointer *int

	variavel3 = 100
	pointer = &variavel3
	fmt.Println(variavel3, pointer) //unreferency


	variavel3 = 150 
	fmt.Println(variavel3, pointer)

}
