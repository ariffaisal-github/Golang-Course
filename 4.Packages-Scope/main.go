package main

import (
	"fmt"

	"example.com/mathlib"
)

func main() { 
	fmt.Println("Calling add on same package")
	fmt.Println(add(3, 4)) // if you call this , must do "go run main.go sum.go"

	fmt.Println("Calling add on different package")
	fmt.Println(mathlib.Multiply(3, 4)) // can run with go run main.go
	// fmt.Println(mathlib.money) // cant use small letter
}

func init() {
	fmt.Println("Running init function which was inside main.go")
}
