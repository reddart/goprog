package main

import "fmt"

func main() {
	age := 50
	if age < 13 {
		fmt.Println("WOW! You are young.")
	} else if age < 20 {
		fmt.Println("Not so young!")
	} else if age < 30 {
		fmt.Println("You 're getting there")
	} else {
		fmt.Println("Out of the world")
	}
}
