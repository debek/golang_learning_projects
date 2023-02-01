package main

import (
	"fmt"
)

var a int

type hotdog int

var b hotdog

func main() {
	a = 42
	b = 43

	//A is int tpye
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	//B is main.hotdog type
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	// Conversion from main.hotdog to int
	a = int(b)
	fmt.Println(a)
	fmt.Printf("%T", a)
}
