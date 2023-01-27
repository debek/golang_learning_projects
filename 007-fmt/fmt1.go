package main

import (
	"fmt"
)

var y = 42

func main() {
	fmt.Println(y)
	fmt.Println("The value is", y)

	fmt.Printf("The value is %v\n", y)
	fmt.Printf("The type is %T\n", y)
	fmt.Printf("The unicode of this number looks like that: %U\n", y)
	fmt.Printf("Printf format is: %#x\t%b\t%x\n", y, y, y)

	//Sprintf - we can assign to variable and then print it
	s := fmt.Sprintf("%#x\t%b\t%x\n", y, y, y)
	fmt.Println("Sprintf format is:", s)
}
