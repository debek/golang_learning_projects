package main

import "fmt"

func main() {
	x := []string{"bond_james", "Shaken", "not stirred", "Martinis, Women"}
	y := []string{"moneypenny_miss", "James Bond", "Literature", "Computer Science"}
	z := []string{"no_dr", "Being evil", "Ice cream", "Sunsets"}

	all := []string{}
	all = append(all, x...)
	all = append(all, y...)
	all = append(all, z...)

	m := make(map[int]string)

	for i, v := range all {
		m[i] = v
		fmt.Println(i, v)
	}
	fmt.Println(m)
}
