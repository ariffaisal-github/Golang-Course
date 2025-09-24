package main

import (
	"fmt"
)

func main() {

	switch p {
	case "hello", "okay":
		fmt.Println("p is either hello or okay")
	case "oka":
		fmt.Println("Noo")
	default:
		fmt.Println("yay")
	}

	if p == "hi" {
		fmt.Println("hi is printed")
	} else {
		fmt.Println("okay is printed")
	}
}

const p string = "okay"
