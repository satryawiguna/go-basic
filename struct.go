package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main() {
	var customerOne Customer

	customerOne.Name = "Satrya Wiguna"
	customerOne.Address = "Jl Kresna Gg I No 1"
	customerOne.Age = 40

	fmt.Println(customerOne)

	customerTwo := Customer{
		Name:    "Erna Widhiastuti",
		Address: "Gg Mawar No 2",
		Age:     38,
	}

	fmt.Println(customerTwo)

	customerThree := Customer{"Maha Sanjaya", "Gg Mawar No 2", 0}

	fmt.Println(customerThree)

	customerOne.sayHello("Kaela")
}
