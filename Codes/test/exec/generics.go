package main

import (
	"fmt"
)

type Person struct {
	height float64
	weight float64
	name   string
}

type Item struct {
	weight float64
	name   string
}
type Storage struct {
	name         string
	location     string
	itemsInStock []Item
}

func PrintAny(any interface{}) {
	switch elem := any.(type) {

	case *Person:
		fmt.Printf("This persons´s name is %s, is %f feet tall and weights %f pounds\n", elem.name, elem.height, elem.weight)

	case *Storage:
		fmt.Printf("This is storage is located in %s, it´s name is %s\n", elem.location, elem.name)
		fmt.Println("Items in stock are:")
		for index, item := range elem.itemsInStock {
			fmt.Printf("%d - The %s weights %f pounds\n", index+1, item.name, item.weight)
		}

	default:
		printHelloWorld()
	}
}
