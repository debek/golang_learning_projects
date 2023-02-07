package main

import "fmt"

func main() {
	slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	slice1 := slice[:5]
	slice2 := slice[5:]
	slice3 := slice[2:7]
	slice4 := slice[1:6]

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)

}
