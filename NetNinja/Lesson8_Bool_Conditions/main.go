package main

import ("fmt")

 
func main()  {
	
	// age := 45

	// fmt.Println(age <= 50)
	// fmt.Println(age >= 45)
	// fmt.Println(age == 50)
	// fmt.Println(age != 50)
	
	// if age < 30 {
	// 	fmt.Println("Age is less than 30")
	// } else if age < 30 {
	// 	fmt.Println(" age is greater than 30")

	// } else  {
	// 	fmt.Println(" age is not less than 45 ")
	// }

	
	names := []string{"mario", "luigi", "yoshu","peach", "edu"}


		for index, value := range names {

			if index == 1 {
				fmt.Println(index)
				continue
			} 
			if index == 4 {
				fmt.Println(" reakinf at pos", value, index)
				break
			}


			fmt.Printf("the vaulue at pos %v is %v", index, value)
		}

	
}