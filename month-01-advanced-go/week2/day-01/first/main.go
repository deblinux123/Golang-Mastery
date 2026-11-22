package main

import "fmt"

type Speaker interface {
	Speek()
}

type Dog struct {
	Name string
}

func (d Dog) Speek() {
	fmt.Println(d.Name, " says woof")
}

type Robot struct {
	Model string
}

func (r Robot) Speek() {
	fmt.Println(r.Model, "says beeb.")
}

func MakeSounc(s Speaker) {
	s.Speek()
}

func main() {
	dog := Dog{Name: "Max"}
	robo := Robot{Model: "R2DF_0F"}

	MakeSounc(dog)
	MakeSounc(robo)
}
