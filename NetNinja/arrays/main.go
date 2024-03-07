package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays and slices")

	var array1 [5]string
	array1[0] = "Posição 1"
	fmt.Println(array1)

	array2 := [5]string{"pos1", "pos2", "pos3", "pos4", "pos5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	slice := []int{10, 11, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice)

	slice = append(slice, 23)
	fmt.Println(slice)

	slice2 := array2[0:3]
	fmt.Println(slice2)

	array2[1] = "Pos changed"
	fmt.Println(slice2) 
	
}
