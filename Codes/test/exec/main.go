package main

func main() {
	// fmt.Println("Hello, World!")
	// fmt.Println(add(5, 6))
	// fmt.Println(fibonacci(6))
	// routine()
	// printHelloWorld()
	// fmt.Println(palindromo("hola"))
	// for {
	// 	var input string
	// 	fmt.Scanln(&input)
	// 	fmt.Println(palindromo(input))
	// }
	// shortVariableDeclaration()
	// for_range()
	// switchExample()
	// fmt.Printf()

	// var arr1 = [5]int{1, 2, 3, 4, 5}
	// Sum(arr1[:])

	// Append_test1()

	// Append_test2()

	//Ptrs_test1()

	//Ptrs_test2()
	//Ptrs_test4(-100, Ptrs_test3())

	//defer-panic-recover test
	// fmt.Printf("Calling test\r\n")
	// test()
	// fmt.Printf("Test completed\r\n")

	person1 := new(Person)
	person2 := new(Person)

	person1.height = 5.6
	person1.weight = 150.6
	person1.name = "John"

	person2.height = 6.2
	person2.weight = 100.9
	person2.name = "Peter"

	storage1 := new(Storage)
	storage2 := new(Storage)

	item1 := new(Item)
	item2 := new(Item)
	item3 := new(Item)
	item4 := new(Item)

	item1.name = "Screwdriver"
	item1.weight = 14.5

	item2.name = "Plastic Box"
	item2.weight = 2.3

	item3.name = "Nails"
	item3.weight = 0.5

	item4.name = "Saw"
	item4.weight = 22.7

	storage1.location = "Alaska"
	storage1.name = "Ralph´s workshop"
	storage1.itemsInStock = []Item{*item1, *item2, *item3}

	storage2.location = "Ohio"
	storage2.name = "The homies junkhouse"
	storage2.itemsInStock = []Item{*item1, *item2, *item4}

	PrintAny(person1)
	PrintAny(person2)

	PrintAny(storage1)
	PrintAny(storage2)
}
