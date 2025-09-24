package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("---------------- NUMERIC TYPES ------------------")
	var int_1 = 234
	fmt.Println(unsafe.Sizeof(int_1))

	var int_2 int64
	int_2 = 123
	fmt.Println(unsafe.Sizeof(int_2))

	var int_3 uint16
	int_3 = 154
	fmt.Println(unsafe.Sizeof(int_3))

	flt_1 := 6.02e230
	fmt.Println(flt_1)
	fmt.Println(unsafe.Sizeof(flt_1))

	var flt_2 float32
	flt_2 = 6.02323425367346352463737743e23
	fmt.Println(flt_2)
	fmt.Println(unsafe.Sizeof(flt_2))

	var complx_1 = 43 - 324i // complex128
	fmt.Println(complx_1)

	var complx_2 complex64
	complx_2 = 23 + 4i
	fmt.Println(complx_2)
	fmt.Println(complex64(complx_1) + complx_2)

	fmt.Println("---------------- BOOLEAN TYPES ------------------")
	var firstBool bool
	firstBool = false
	fmt.Println(firstBool)
	fmt.Println("---------------- STRING TYPES ------------------")
	var name string
	name = "Arif"
	fmt.Println(unsafe.Sizeof(name))
	fmt.Println(string(name[0]))
}
