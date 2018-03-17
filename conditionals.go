package main

import (
	"fmt"
)

func main() {
	var n uint64

	fmt.Scan(&n)

	if n%2 != 0 {
		fmt.Println("Weird")
	} else {
		if n >= 6 && n <= 20 {
			fmt.Println("Weird")
		} else {
			fmt.Println("Not Weird")
		}
	}
}
