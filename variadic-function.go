package main

import "fmt"

func sumAll(numbers ...int) (result int) {
	result = 0
	for _, number := range numbers {
		result = result + number
	}
	return
}

func main() {
	total := sumAll(10, 10, 10)
	fmt.Println("Total : ", total)

	// mengirimkan data yang sudah menjadi array ke varargs
	numbers := []int{10, 10, 10, 10}
	fmt.Println("Total : ", sumAll(numbers...))
}
