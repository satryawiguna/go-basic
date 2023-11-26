package main

import "fmt"

func main() {
	var data [3]string

	data[0] = "Satrya"
	data[1] = "Erna"
	data[2] = "Bejo"

	fmt.Println(data)

	var dataKedua = [3]int{
		10,
		20,
		30,
	}

	fmt.Println(dataKedua)

	var dataKetiga = [...]string{
		"Anjing",
		"Kurap",
		"Kudis",
		"Koreng",
	}

	fmt.Println(len(dataKetiga))
}
