package main

import "fmt"

func main() {
	name := "John"
	age := 25
	// fmt.Printf - display text on the screen
	fmt.Printf("My name is %s and I am %d years old\n", name, age)
	// fmt.Sprintf - return text as a value
	message := fmt.Sprintf("My name is %s and I am %d years old", name, age)
	fmt.Println(message)
}
