package main

func suma2(a int, b int, c chan int) {
	suma := a + b
	c <- suma
}
