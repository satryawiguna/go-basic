package main

import "fmt"

func main() {
	var title string = "Lika-liku perselingkuhan"

	switch title {
	case "Maju kena, mundur kena":
		fmt.Println("Bernilai salah")

	case "Hentai":
		fmt.Println("Bernilai benar")

	default:
		fmt.Println("Bernilai salah")
	}

	switch titleLength := len(title); titleLength > 5 {
	case true:
		fmt.Println("Bernilai benar")
	case false:
		fmt.Println("Bernilai salah")
	}

}
