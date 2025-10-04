package main

import "fmt"

func add(a int, b int) int { // add two numbers
	return a + b
}

func init() {
	fmt.Println("Running init function which was inside sum.go")
}

