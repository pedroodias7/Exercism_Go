package main


import "fmt"



func main()  {

	age := 34
	name := "Pedro"



	fmt.Print("Hello \n")
	fmt.Print("World")
	fmt.Print("new line \n")


	fmt.Println("hello pedro!")
	fmt.Println("Adeus Pedro!")


	fmt.Println("My age is", age, "and my name is", name,)


	// Print (formated sttring) %_ = format specifier

	fmt.Printf("My age is %v and my name is %v \n", age, name)
	fmt.Printf("My age is %q and my name is %q\n", age, name)
	fmt.Printf("age is of type %T\n", age,)
	fmt.Printf("you score %0.2f points \n", 24.123)

	  

	// Sprinf (save formatted strings )


	var str = fmt.Sprintf("My age is %v and my name is %v \n", age, name)
	fmt.Println("the saved string is :", str)


}