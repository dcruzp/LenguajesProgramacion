package main

import "fmt"

func for_range() {
	//for-each loop
	a := []string{"C++", "C", "Python", "Go"}
	for i, j := range a {
		fmt.Println("index: ", i, " ", "element: ", j)
	}
}
