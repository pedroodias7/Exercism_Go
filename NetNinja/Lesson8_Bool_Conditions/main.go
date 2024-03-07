package main

import ("fmt")

 
func main()  {
	
	idade := 13

	// fmt.Println(age <= 50)
	// fmt.Println(age >= 45)
	// fmt.Println(age == 50)
	// fmt.Println(age != 50)
	
	if idade < 30 {
		fmt.Println("A idade é menor que 30")
	} else if idade > 30 {
		fmt.Println("A idade é superior a 30")

	} else  {
		fmt.Println(" nao sei, faz tu ")
	}

	
	// names := []string{"mario", "luigi", "yoshu","peach", "edu"}


	// 	for index, value := range names {

	// 		if index == 1 {
	// 			fmt.Println(index)
	// 			continue
	// 		} 
	// 		if index == 4 {
	// 			fmt.Println(" reakinf at pos", value, index)
	// 			break
	// 		}


	// 		fmt.Printf("the vaulue at pos %v is %v", index, value)
	// 	}

	
}