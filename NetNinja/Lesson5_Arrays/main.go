package main


import "fmt"

func main()  {
	var ages [3]int = [3]int{20, 25, 30} 

	names := [4]string{"yoshi", "mario", "apple", "yellow"}

	names[1] = "Luigi"

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// slices (arrays under the hood)

	var score = []int{100, 50, 60}

	score[2] = 22

	score = append(score, 85)

	fmt.Println(score, len(score))


	// slice ranges

	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]

	fmt.Println(rangeOne)
	fmt.Println(rangeTwo)
	fmt.Println(rangeOne, rangeTwo, rangeThree)

	rangeOne = append(rangeOne, "koopa")
	fmt.Println(rangeOne)

}