package main

import "fmt"

type Person struct {
	Name, Address string
	Age           int
}

type HasName interface {
	GetName() string
}

func SayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

func (person Person) GetName() string {
	return person.Name
}

func main() {
	person := Person{
		Name:    "Ganjar Pranowo",
		Address: "Jl. Danau Teme No 123",
		Age:     123,
	}

	SayHello(person)
}
