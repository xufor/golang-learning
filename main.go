package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	var name string = "Akash"
	fmt.Printf("This is my name %s\n", name)
	age := 25
	fmt.Printf("My age is %d\n", age)
	var city, country string = "G. Noida", "India"
	var (
		isEmployed bool   = true
		sal        int    = 60000
		position   string = "Junior Developer"
	)
	fmt.Printf("My city is %s, and country is %s", city, country)
	fmt.Printf("I am %s.")
}
