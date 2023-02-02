package main

import "fmt"

type lol int

var i lol

func main() {

	fmt.Println(i)
	fmt.Printf("%T\n", i)

	i = 4
	fmt.Println(i)
}
