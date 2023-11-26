package main

import "fmt"

func main() {
	var nilaiPertama int64 = 1
	var nilaiKedua int32 = int32(nilaiPertama)

	fmt.Println(nilaiKedua)

	var kalimat string = "Nama Saya adalah Satrya"
	var e = kalimat[1]
	eString := string(e)

	fmt.Println(e)
	fmt.Println(eString)
}
