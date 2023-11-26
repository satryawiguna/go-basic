package main

import "fmt"

func main() {
	var nilaiPertama int16 = 32767

	var nilaiKedua int32 = int32(nilaiPertama) + 2

	fmt.Println(nilaiKedua)

	namaOrangPertama := "Satrya"
	var nameOrangKedua string = "Satrya Wiguna"
	var result bool = len(namaOrangPertama) < len(nameOrangKedua)

	fmt.Println(result)
}
