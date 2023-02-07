package main

import (
	"fmt"
)

func main() {
	jb := []string{"James", "Bond", "Chocolate", "martini"}
	//fmt.Println(jb)
	mp := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnut"}
	//fmt.Println(mp)

	xp := [][]string{jb, mp}
	//fmt.Println(xp)

	//fmt.Println(xp[0][0])

	//Method 1
	for x := 0; x <= len(xp)-1; x++ {
		fmt.Println(x)
		fmt.Println(xp[x])
		for y := 0; y <= len(xp[x])-1; y++ {
			fmt.Println(xp[x][y])
		}
	}

	//Method 2. Cam be also without _
	for z, _ := range xp {
		fmt.Println(xp[z])
		for y, _ := range xp[z] {
			fmt.Println(xp[z][y])
		}
	}

}
