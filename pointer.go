package main

import "fmt"

func main() {
	var nilaiA = [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	nilaiB := &nilaiA

	fmt.Println(nilaiB)

	nilaiB[0] = "Libur"

	fmt.Println(nilaiA)
}
