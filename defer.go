package main

import "fmt"

func main() {
	pertama()
}

func pertama() {
	defer kedua()

	fmt.Println("Dijalankan Pertama")
}

func kedua() {
	fmt.Println("Dijalankan Kedua")
}
