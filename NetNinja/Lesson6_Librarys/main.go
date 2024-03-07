package main

import (
	"fmt"
	"sort"
)

func main()  {
	
	// greeting := "Hello there friends"

	// // fmt.Println(strings.Contains(greeting, "Hello")) // return true
	// // fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	// // fmt.Println(strings.ToUpper(greeting))
	// // fmt.Println(strings.Index(greeting, "th"))
	// fmt.Println(strings.Split(greeting, ""))


	ages := []int{45, 20, 35, 75, 60, 50, 25}

	sort.Ints(ages)
	fmt.Println(ages)


	index := sort.SearchInts(ages, 30)
	fmt.Println(index)



	names := []string{"yoshi", "mario", "apple", "yellow"}

	sort.Strings(names)
	fmt.Println(names)


	fmt.Println(sort.SearchStrings(names, "mario"))

}