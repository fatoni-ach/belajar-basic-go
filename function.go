package main

import "fmt"

func penjumlahan(a int, b int) int {
	return a + b
}

func getFullName() (string, string) {
	firstName := "Achmad"
	lastName := "Fatoni"
	return firstName, lastName
}

// Named Return Value
func getFullNameWithNamedValue() (firstname string, lastname string, age int) {
	firstname = "Achmad"
	lastname = "fatoni"
	age = 25
	return
}

func main() {
	a := 10
	b := 20
	fmt.Println("Hasil Penjumlahan ", a, " + ", b, " = ", penjumlahan(a, b))

	firstName, _ := getFullName()
	fmt.Println("Nama lengkap saya adalah : ", firstName)

	firstName, lastname, age := getFullNameWithNamedValue()
	fmt.Println("Nama lengkap saya adalah : ", firstName, lastname)
	fmt.Println("Umur : ", age)
}
