package main

import "fmt"

type lol int

var x lol

func main() {

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 4
	fmt.Println(x)
}
