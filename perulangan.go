package main

import "fmt"

func main() {
	fmt.Println("Perulangan : ")

	slice := []string{"ikan", "daging", "mie", "telor"}

	// Standar For
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// Using For Range
	for _, value := range slice {
		fmt.Println(value)
	}

	// Using Map
	dataMap := make(map[string]string)
	dataMap["name"] = "Fatoni"
	dataMap["Address"] = "Bangkalan"
	dataMap["Job Title"] = "Software Engineer"

	for key, value := range dataMap {
		fmt.Println(key, " = ", value)
	}
}
