package main

import "fmt"

func main() {
	var i int8 = 0

	for i <= 10 {
		fmt.Println("Perulangan ke", i)
		i++
	}

	for j := 0; j <= 10; j++ {
		fmt.Println("Perulangan ke", j)
	}

	books := map[string]string{
		"title":       "Judul pertama",
		"author":      "Satrya Wiguna",
		"description": "Description judul pertama",
	}

	for index, value := range books {
		fmt.Println("Perulangan index:", index, "bernilai:", value)
	}

	for _, value := range books {
		fmt.Println("Perulangan bernilai:", value)
	}

	for _, value := range books {
		if value == "Satrya Wiguna" {
			break
		}

		fmt.Println("Perulangan bernilai:", value)
	}

	for _, value := range books {
		if value == "Satrya Wiguna" {
			continue
		}

		fmt.Println("Perulangan bernilai:", value)
	}
}
