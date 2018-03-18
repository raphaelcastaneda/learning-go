package main

import (
	"fmt"
)

func to_binary_list(n int) []int {
	var remainder int
	result := make([]int, 0)
	dividend := n
	for dividend > 0 {
		remainder = dividend % 2
		result = append(result, remainder)
		dividend = dividend / 2
	}
	return result
}

func largest_consecutive_ones(a []int) int {
	longest_string := 0
	current_string := 0
	for _, n := range a {
		if n == 1 {
			current_string += 1
			if current_string > longest_string {
				longest_string = current_string
			}
		} else {
			current_string = 0
		}
	}
	return longest_string
}

func main() {
	var n int
	fmt.Scan(&n)
	b := to_binary_list(n)
	fmt.Println(largest_consecutive_ones(b))
}
