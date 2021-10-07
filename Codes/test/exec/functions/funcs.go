package main

import "fmt"

// multiple results
func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// functions continued
func add(x, y int) int {
	return x + y
}

func main() {
	a, b := swap("hello", "world")
	// underscore
	fmt.Println(a, b)
	fmt.Println(add(5, 6))
	fmt.Println(split(100))
}
