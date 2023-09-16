package main

import "fmt"

func main() {
	// Declaration
	var kata string = "Hello World"
	var array [3]string
	months := [...]string{"Jan", "Feb", "March", "Apr", "May", "Jun", "Jul", "August", "Sept", "Okt", "Nov", "Des"}

	array[0] = "Hai.."
	array[1] = "Nama saya Toni"
	array[2] = "Terima kasih"

	// Prosess
	fmt.Println(kata)

	var a int = 10

	fmt.Println(a)

	a += 10

	fmt.Println(a)

	if false {
		for i := 0; i < len(array); i++ {
			fmt.Println(array[i])
		}
	}

	if false {
		var b = array[:2]

		for i := 0; i < len(b); i++ {
			fmt.Println(b[i])
		}
	}

	if true {
		for i := 0; i < len(months); i++ {
			fmt.Println(months[i])
		}
	}
}
