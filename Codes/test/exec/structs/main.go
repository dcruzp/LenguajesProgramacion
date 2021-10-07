package main

type Animal struct {
	name  string
	class string
}

func main() {

	perro := new(Animal)
	perro.name = "Simba"
	perro.class = "Mamifero"

}
