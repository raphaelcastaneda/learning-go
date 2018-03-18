package main

import (
	"fmt"
)

func sum_slice(a []int) {
	sum := 0
	for _, n := range a {
		sum += n
	}
	return sum
}
func sum_hourglass(a [][]int, i, j int) int {
	sum := 0
	sum += sum_slice(a[i][j : j+3])
	sum += a[i+1][j+1]
	sum += sum_slice(a[i+2][j : j+3])
	return sum
}

func main() {
	var n int
	var largest_sum int

	// Build the 2D array
	a := [6][6]int{}
	for i, _ := range a {
		for j, _ := range a[i] {
			fmt.Scan(&n)
			a[i][j] = n
		}
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			sum := sum_hourglass(a, i, j)
			if (i == 0 && j == 0) || sum > largest_sum {
				largest_sum = sum
			}
		}
	}
	fmt.Println(largest_sum)
}
