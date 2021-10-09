package main

import "fmt"

type Country struct {
	name       string
	capital    string
	language   string
	population int
}

type Person struct {
	name     string
	lastname string
	age      int
}

func (p *Person) Fullname() string {
	return p.name + " " + p.lastname
}

type Point struct {
	x int
	y int
}

func NewPoint(newX, newY int) *Point {
	return &Point{x: newX, y: newY}
}

// func NewCity(name string, population int) *City {
// 	if name == "" {
// 		return nil
// 	}
// 	return &City{
// 		name,
// 		population}
// }
func main() {

	// ru := new(Country) //pointer to struct type using new
	// ru.name = "Russia"
	// ru.capital = "Moscow"
	// ru.language = "russian"
	// ru.population = 146171015

	person1 := new(Person)
	person1.name = "John"
	person1.lastname = "Smith"
	person1.age = 30

	fmt.Println(person1.Fullname())

	// p1 := NewPoint(1, 2)

	// sp := Country{
	// 	"Spain",
	// 	"Madrid",
	// 	"spanish",
	// 	47450795} // struct type

	// var it *Country = &Country{
	// 	name:       "Italy",
	// 	capital:    "Rome",
	// 	language:   "italian",
	// 	population: 60317116}

	// fmt.Println(sp)
	// fmt.Println(*ru)
	// fmt.Printf("Name: %s\n", ru.name)
	// fmt.Printf("Capital: %s\n", ru.capital)
	// fmt.Printf("Language: %s\n", ru.language)
	// fmt.Printf("Population: %d\n", ru.population)

	// fmt.Println(it.name)
	// fmt.Println(*it)
	// fmt.Println(p1)
}
