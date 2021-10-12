package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 2; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func suma(a int, b int) {
	time.Sleep(100 * time.Millisecond)
	suma := a + b
	fmt.Println(suma)
}
func main() {
	// go say("world")
	// go suma(1, 2)
	// say("hello")
	c := make(chan int)
	go suma2(10, 20, c)
	x := <-c
	fmt.Println(x)

}
