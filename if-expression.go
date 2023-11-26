package main

import "fmt"

func main() {
	var name string = "Satrya Wiguna"

	if name == "Satrya" {
		fmt.Println("Bernilai Benar")
	}

	nameDua := "Ganjar Pranowo"

	if nameDua == "Prabowo" {
		fmt.Println("Bernilai Salah")
	} else {
		fmt.Println("Bernilai Benar")
	}

	if nameDua == "Prabowo" {
		fmt.Println("Bernilai Salah")
	} else if nameDua == "Ganjar Pranowo" {
		fmt.Println("Bernilai Benar")
	} else {
		fmt.Println("Bernilai Salah")
	}

	if nameLength := len(name); nameLength > 5 {
		fmt.Println("Bernilai Benar")
	} else {
		fmt.Println("Bernilai Salah")
	}
}
