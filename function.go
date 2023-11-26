package main

import "fmt"

func main() {
	sayHello()

	sayHelloTo("Satrya", "Wiguna")

	var sayHelloToWithReturn string = sayHelloToWithReturn("Satrya Wiguna")

	fmt.Println(sayHelloToWithReturn)

	firstName, lastName := sayHelloToWithMultipleReturn()

	fmt.Println("Hello", firstName, lastName)

	firstNamDua, _ := sayHelloToWithMultipleReturn()

	fmt.Println("Hello", firstNamDua)
}

func sayHello() {
	fmt.Println("Welcome to Hellomart, Happy shopping...!")
}

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName, "Welcome to Hellomart, Happy shopping...!")
}

func sayHelloToWithReturn(name string) string {
	return "Hello " + name
}

func sayHelloToWithMultipleReturn() (string, string) {
	return "Erna", "Widhiastuti"
}
