package main

import "fmt"

func main() {
	runApp(true)

	fmt.Println("Tidak terjadi ERROR")
}

func runApp(e bool) {
	defer endApp()

	if e {
		panic("ERROR")
	}
}

func endApp() {
	fmt.Println("Dijalankan Terakhir")

	message := recover()
	fmt.Println("Terjadi Panic:", message)
}
