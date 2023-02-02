package main

import (
	"fmt"
)

const (
	j int     = 42
	k float64 = 42.78
	l string  = "James Bond"
)

func main() {
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)
	fmt.Printf("%T\n", j)
	fmt.Printf("%T\n", k)
	fmt.Printf("%T\n", l)
}
