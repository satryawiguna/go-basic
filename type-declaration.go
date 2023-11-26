package main

import "fmt"

func main() {
	type NoPlat string

	var noPlat NoPlat = "DK 6996 KO"

	fmt.Println(noPlat)
}
