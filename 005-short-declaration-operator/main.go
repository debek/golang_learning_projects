package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	//New value of declared variable before
	x = 99
	fmt.Println(x)

	//Expression 100 + 24 is making all line as statement
	y := 100 + 24
	fmt.Println(y)
}
