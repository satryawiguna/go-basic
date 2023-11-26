package main

import "fmt"

func Random() any {
	return "OK"
}

func main() {
	result := Random()
	resultOne := result.(string)

	fmt.Println(resultOne)
}
