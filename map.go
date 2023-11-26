package main

import (
	"fmt"
)

func main() {
	var dataMap = map[string]string{
		"name":    "Satrya Wiguna",
		"address": "Jl. Kresna No 2",
		"city":    "Gianyar",
		"poscode": "12345",
	}

	fmt.Println(dataMap)

	book := make(map[string]string)

	book["title"] = "Harry Porter"
	book["author"] = "Richard Sucherberg"

	fmt.Println(book)

	book["description"] = "There is no description"

	fmt.Println(book)

	delete(book, "description")

	fmt.Println(book)
}
