package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHi(name string) {
	fmt.Println("Hello ", name, "My Name is ", customer.Name)

}

func main() {
	var eko Customer
	eko.Name = "eko"
	eko.Address = "Indonesia"
	eko.Age = 30

	eko.sayHi("Joko")
}
