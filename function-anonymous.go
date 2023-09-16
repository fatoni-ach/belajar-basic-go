package main

import "fmt"

type Filter func(string) string

func sayHelloWithFunction(name string, filter Filter) {
	fmt.Println("Hello ", filter(name))
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	}

	return name
}

func main() {
	sayHelloWithFunction("anjing", spamFilter)
	sayHelloWithFunction("toni", spamFilter)

	// Anonymous function
	checkAdmin := func(name string) string {
		if name == "admin" {
			return "!!__SuperUser__!!"
		}
		return name
	}
	sayHelloWithFunction("admin", checkAdmin)
	sayHelloWithFunction("staff", checkAdmin)
}
