package main

import "fmt"

func main() {
	title, author, price := book()

	fmt.Println(title, author, price)
}

func book() (title string, author string, price int32) {
	title = "Janda Perawan"
	author = "Choirul Gaban"
	price = 200000

	return title, author, price
}
