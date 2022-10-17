package main

import (
	"fmt"
)

func Testing_arr() {
	var coll [10]int

	//Conditional styled for loop
	// for i := 0; i < len(coll); i++ {
	// 	fmt.Printf("Array coll at index %d is %d\n", i, coll[i])
	// }

	//Range styled for loop
	for i := range coll {
		fmt.Printf("Array (Aquí probando, probando) coll at index %d is %d\n", i, coll[i])
	}
}

//Value type array
func Arr_managment1() {
	var arr1 [3]int
	arr2 := arr1

	arr2[0] = 1

	fmt.Printf("arr1 at index 0 is %d\n", arr1[0])
	fmt.Printf("arr2 at index 0 is %d\n", arr2[0])
}

//Ref Type array
func Arr_managment2() {
	var arr1 [3]int
	arr2 := &arr1

	arr2[0] = 1

	fmt.Printf("arr1 at index 0 is %d\n", arr1[0])
	fmt.Printf("arr2 at index 0 is %d\n", arr2[0])
}

//Empty return
func Sum_empty_return(a *[3]float64) (sum float64) {
	for _, v := range a {
		sum += v
	}
	return
}

//Passing slices to functions
func Sum(elems []int) {
	s := 0
	for i := range elems {
		s += elems[i]
	}

	fmt.Printf("Sum is %d", s)
}

//Append
func Append_test1() {
	sli1 := []int{1, 2, 3}
	sli1 = append(sli1, 4, 5, 6)
	fmt.Println(sli1)
}

func Append_test2() {
	sli1 := []int{1, 2, 3}
	sli2 := []int{4, 5, 6}
	sli3 := append(sli1, sli2...)
	fmt.Println(sli3)
}
