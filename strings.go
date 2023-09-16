package main

import (
	"fmt"
	"strings"
)

func main() {
	word := "Hallo semuanya"

	fmt.Println("word : ", word)
	fmt.Println("word UpperCase : ", strings.ToUpper(word))
	fmt.Println("word LowerCase : ", strings.ToLower(word))
}
