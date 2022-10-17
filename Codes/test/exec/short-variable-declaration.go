package main

import "fmt"

func shortVariableDeclaration() {
	sum := 0
	for i := 10; i >= 0; i-- {
		sum -= i
	}

	fmt.Println(sum)

}
