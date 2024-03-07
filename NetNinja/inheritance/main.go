package main

import "fmt"


type pessoa struct {
	nome string
	idade int8
	altura int8	
}

type estudante struct{
	pessoa
	curso string
	faculdade string
}

func main()  {
	fmt.Println("Herança")
	p1 := pessoa{"Pedro", 20, 8}
	fmt.Println(p1)


	e1 := estudante{p1, "inf", "istec"}
	fmt.Println(e1.idade)
}