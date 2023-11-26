package main

import "fmt"

func main() {
	var numbers = []int{10, 20, 30, 40, 50, 60}

	var total int = sumAll(numbers...)

	fmt.Println(total)

	totalDua := sumAll(10, 20, 30, 40, 50, 60)

	fmt.Println(totalDua)
}

func sumAll(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total = total + number
	}

	return total
}
