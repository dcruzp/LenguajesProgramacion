package main

import "fmt"

type Recibir interface {
	RecibirClase()
}

type Impartir interface {
	ImpartirClase()
}

type Cobrar interface {
	CobrarSalario()
}

type Persona struct {
	name string
}
type Estudiante struct {
	Persona
}

type Profesor struct {
	Trabajador
}

type AlumnoAyudante struct {
	Estudiante
}

type Trabajador struct {
	Persona
}

func (p Profesor) ImpartirClase() {
	fmt.Printf("El profesor %s imparte la clase\n", p.name)
}

func (e Estudiante) RecibirClase() {
	fmt.Printf("El estudiante %s recibe la clase\n", e.name)
}
func (t Trabajador) CobrarSalario() {
	fmt.Printf("El trabajador %s cobra el salario\n", t.name)
}

func (a AlumnoAyudante) ImpartirClase() {
	fmt.Printf("El alumno %s  imparte la clase\n", a.name)
}

func main() {
	fernan := new(Profesor)
	fernan.name = "Fernando"
	david := new(Estudiante)
	david.name = "David"
	ale := new(AlumnoAyudante)
	ale.name = "Alejandro"
	julio := new(Trabajador)
	julio.name = "Julio"

	fernan.CobrarSalario()
	fernan.ImpartirClase()

	david.RecibirClase()

	ale.RecibirClase()
	ale.ImpartirClase()

	julio.CobrarSalario()

}
