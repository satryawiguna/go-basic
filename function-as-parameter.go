package main

import "fmt"

type Filter func(string) string

func main() {
	sayHelloWithFilter("Wiguna", filterWord)
	sayHelloWithTypeDeclarationFilter("Satrya", filterWord)
}

func sayHelloWithFilter(name string, filter func(string) string) {
	fmt.Println("Hello " + filter(name))
}

func sayHelloWithTypeDeclarationFilter(name string, filter Filter) {
	fmt.Println("Hello " + filter(name))
}

func filterWord(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}
