package main

import (
	"fmt"
	"time"
)

func main() {
	date := time.Now().Year()

	for i := 1989; i <= date; i++ {
		fmt.Println(i)
	}

	bd := 1989
	for bd <= date {
		fmt.Println(bd)
		bd++
	}

}
