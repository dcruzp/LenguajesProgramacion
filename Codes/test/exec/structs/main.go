package main

import "fmt"

type Country struct {
	name       string
	capital    string
	language   string
	population int
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

	ru := new(Country) //pointer to struct type using new
	ru.name = "Russia"
	ru.capital = "Moscow"
	ru.language = "russian"
	ru.population = 146171015

	sp := Country{
		"Spain",
		"Madrid",
		"spanish",
		47450795} // struct type

	var it *Country = &Country{
		"Italy",
		"Rome",
		"italian",
		60317116} //pointer to struct type

	fmt.Println(sp)
	fmt.Println(*ru)
	fmt.Println(ru)
	fmt.Println(it.name)
	fmt.Println(*it)
}
