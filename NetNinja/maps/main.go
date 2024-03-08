package main

import "fmt"

func main() {
	fmt.Println("Maps")

	user := map[string]string {

		"nome": 	"Pedro",
		"sobrenome": 	"Dias",
	}

	fmt.Println(user["nome"])

	user2 := map[string]map[string]string {
		"nome":{
			"primeiro": "Pedro",
			"ultimo": "Dias",
		},
		"curso":{
			"nome": "Informatica",
			"Loc": "Porto",
		},
	}


	fmt.Println(user2)
	delete(user2, "nome")
	fmt.Println(user2)

}