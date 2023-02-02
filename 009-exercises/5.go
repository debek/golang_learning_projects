package main

import "fmt"

type lol2 int

var x lol2

var y int

func main() {

	x = 4

	fmt.Println(x)
	fmt.Printf("%T\n\n", x)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T", y)
}
