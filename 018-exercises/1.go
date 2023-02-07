package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}

	fmt.Println(array)

	for _, v := range array {
		println(v)
	}

	fmt.Printf("Array type is %T", array)
}
