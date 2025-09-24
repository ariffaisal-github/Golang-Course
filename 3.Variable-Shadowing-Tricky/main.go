package main

import (
	"fmt"
)

func sum(p int, q int) int {
	sum := p + q
	return sum
}

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)

	// ------------------------ Shadowing a function with a variable
	sum := sum(a, b)
	fmt.Println("Sum = ", sum)
	// the following give error now , cuz from now on , sum is a variable
	// var sum int
	// sum = sum(a, b)
	// fmt.Println("Sum = ", sum)
	// sum2 := sum(b, a)
	// fmt.Println("Sum = ", sum2)

	// ---------------------  Shadowing with short declarations in loops
	x := 1
	for i := 0; i < 3; i++ {
		x := x + i // shadows outer x each time
		fmt.Println(x)
	}
	fmt.Println("outer:", x) // still 1

	main := 123
	fmt.Println(main)

	// ---------------------- Shadowing with err (the most common bug in Go codebases!)
	err := fmt.Errorf("Outer Error")
	if err != nil {
		fmt.Println("error: ", err)
	}

	if err := fmt.Errorf("Inside If Error"); err != nil {
		fmt.Println("error again:", err)
	}
	// At this point, outer err is unchanged
	// inner err only lived inside the if block
	fmt.Println("err after if:", err)
}
