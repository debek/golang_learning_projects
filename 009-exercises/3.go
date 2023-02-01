package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	fmt.Printf("x type is: %T\n", x)
	fmt.Printf("x value is: %v\n\n", x)

	fmt.Printf("y type is: %T\n", y)
	fmt.Printf("y value is: %v\n\n", y)

	fmt.Printf("z type is: %T\n", z)
	fmt.Printf("z value is: %v\n\n", z)

	output := fmt.Sprintf("%v\t%v\t%v\t", x, y, z)
	println(output)
}
