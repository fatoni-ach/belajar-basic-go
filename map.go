package main

import "fmt"

func main() {
	fmt.Println("Type data Map : ")

	identity := map[string]string{
		"Name":    "Fatoni",
		"Address": "Bangkalan",
	}

	identity["Job Title"] = "Software Engineer"
	fmt.Println(identity)
	fmt.Println("Panjang Map : ", len(identity))

}
