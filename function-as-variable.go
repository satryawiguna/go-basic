package main

import "fmt"

func main() {
	goodBye := getGoodBye

	fmt.Println(goodBye("Satrya"))
}

func getGoodBye(name string) string {
	return "Good bye " + name
}
