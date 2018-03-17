package main

import (
	"fmt"
)

func PrintReverse(n int, arr []int) {
	// Print a space separated string with each element in arr in reverse order

	for i := n - 1; i >= 0; i-- {
		if i != n-1 {
			fmt.Printf(" ")
		}
		fmt.Printf("%d", arr[i])
	}
	fmt.Printf("\n")

}

func main() {

	var length int
	fmt.Scan(&length)

	var arr = make([]int, length) // Using a slice here t
	var tmp int

	for i := 0; i < length; i++ {
		fmt.Scan(&tmp)
		arr[i] = tmp
	}

	PrintReverse(length, arr)

}
