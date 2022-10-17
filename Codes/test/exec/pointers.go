package main

import (
	"fmt"
)

func Ptrs_test1() {
	number := 10
	fmt.Printf("The value of a is %d, and its memory address is %d", number, &number)
}

func Ptrs_test2() {
	number := 10
	fmt.Printf("The value of a is %d, and its memory address is %d \n", number, &number)

	var intPtr *int = &number
	fmt.Printf("The value stored where intPtr points to is %d and intPtr points to %d", *intPtr, intPtr)
}

func Ptrs_test3() *int {
	a := 25
	return &a
}

func Ptrs_test4(number int, ptr *int) {
	fmt.Printf("ptr value is %d", *ptr)
}
