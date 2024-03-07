package main

import "fmt"


type user struct {
	nome string
	idade int8
	endereco endereco
}


type endereco struct {
	morada string
	codPostal int32
}


func main(){
	fmt.Println("Arquivo structs")


	var u user
	u.nome = "Pedro"
	u.idade = 24
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua da Boavista", 0}

	u2 := user{"Pedro", 21, enderecoExemplo}
	fmt.Println(u2)

	u3 := user{nome: "Pedro"}
	fmt.Println(u3)
}