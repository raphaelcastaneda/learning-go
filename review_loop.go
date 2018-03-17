package main

import (
	"fmt"
)

func split_string(s string) (string, string) {
	var evens string
	var odds string

	for i, c := range s {
		if i%2 == 0 {
			evens = evens + string(c)
		} else {
			odds = odds + string(c)
		}
	}
	return evens, odds
}

func main() {

	var c int
	var s string

	fmt.Scan(&c)

	for i := 0; i < c; i++ {
		fmt.Scan(&s)
		evens, odds := split_string(s)
		fmt.Println(evens, odds)
	}

}
