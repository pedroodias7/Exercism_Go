package main

import "fmt"

func main() {

	returno := func(text string)string {
		return fmt.Sprintf("Recebido - %s", text)
	}("Passando parametro")

	fmt.Println(returno)
}