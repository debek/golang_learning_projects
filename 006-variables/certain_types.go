package main

import (
	"fmt"
)

var (
	c string
	d int
)

func main() {
	// DECLARE a variable to be of a certain TYPE
	// and then ASSIGN a VALUE of that type to the variable

	fmt.Println("printing the value of c", c, "ending")
	fmt.Printf("%T\n", c)

	c = "Bond, James Bond"

	fmt.Println(c)
	fmt.Printf("%T\n", c)

	fmt.Println(d)
	fmt.Printf("%T", d)
}
