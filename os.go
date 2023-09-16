package main

import (
	"fmt"
	"os"
)

func main() {
	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSWORD")

	fmt.Println("Username : ", username)
	fmt.Println("Password : ", password)
}
