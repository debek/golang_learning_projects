package main

import (
	"fmt"
)

func main() {
	s := "Hello, playground"
	fmt.Println(s)
	fmt.Printf("%s\n", s)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	//This below makes new line after loop
	fmt.Println("")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%d ", s[i])
	}
	//This below makes new line after loop
	fmt.Println("")
	for i, v := range s {
		fmt.Printf("%#U \t %d", v, i)
	}
}
