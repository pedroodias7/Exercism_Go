package main

import (
	"fmt"
	"math"
)

func sayGrettings(n string) {
	fmt.Printf("Good morning %v \n", n)
}

func sayBye(s string) {
	fmt.Printf("Good bye %v \n", s)
}

func cycleNames(n []string, f func(string)) {

	for _, v := range n {
		f(v)
	}
}

func cycleArea(r float64) float64 {
	return math.Pi * r * r
}

func calculo(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	sayGrettings("mario")

	sayGrettings("luigi")

	sayBye("edu")

	cycleNames([]string{"cloud", "tifa", "barrot"}, sayGrettings)
	cycleNames([]string{"cloud", "tifa", "barrot"}, sayBye)

	a1 := cycleArea(10.5)
	a2 := cycleArea(15)

	fmt.Println(a1, a2)
	fmt.Printf("Circle 1 is %0.3f and circle is %0.3f \n", a1, a2)

	resultadosSoma, _ := calculo(10, 15)
	fmt.Println(resultadosSoma)

}
