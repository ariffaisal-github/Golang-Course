package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func init() {
	fmt.Println("Running init function which was inside sum.go")
}
