package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "Host Database")
	username := flag.String("username", "root", "Username Database")
	password := flag.String("password", "root", "Password Database")

	flag.Parse()

	fmt.Println("Host : ", *host)
	fmt.Println("Username : ", *username)
	fmt.Println("Password : ", *password)

	// go run flag.go -host=localhost -username=toni -password=password
}
