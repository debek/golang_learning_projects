package main

import "fmt"

func main() {
	n, err := fmt.Println("Hello", 42, true)
	fmt.Println(n)
	fmt.Println(err)

	print_Me()
}

func print_Me() {
	fmt.Println("Print me")
}
