package main

import ("fmt")

 
func main()  {
	// x := 0

	// for  x < 5 {
	// 	fmt.Println("Value is x", x)
	// 	x++

	// }

	// for i := 0; i < 5; i++ {
	// 	fmt.Println("Value is x", x)
	// 	x++
		
	// }


	names := []string{"mario", "luigi", "yoshu","peach"}


	// for i :=; i < len(names); i++ {

	// 	fmt.Println(names[i])
	// }

	for index, value := range names {
		fmt.Println("the position in", index, value)
		value = "new string"
	}

	for _, value := range names {
		fmt.Println("the position in", value)
	}

	fmt.Println(names)	 
	
}