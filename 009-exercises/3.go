package main

import "fmt"

var g int = 42
var h string = "James Bond"
var z bool = true

func main() {
	fmt.Printf("g type is: %T\n", g)
	fmt.Printf("g value is: %v\n\n", g)

	fmt.Printf("h type is: %T\n", h)
	fmt.Printf("h value is: %v\n\n", h)

	fmt.Printf("z type is: %T\n", z)
	fmt.Printf("z value is: %v\n\n", z)

	output := fmt.Sprintf("%v\t%v\t%v\t", g, h, z)
	println(output)
}
