package main

import (
	"fmt"
)

func main() {

	var entry_count int
	var name string
	var phone_number string
	var query string

	phone_book := make(map[string]string)

	fmt.Scan(&entry_count)

	for i := 0; i < entry_count; i++ {
		fmt.Scanf("%s %s", &name, &phone_number)
		phone_book[name] = phone_number
	}

	for {
		num, err := fmt.Scanln(&query)
		if num == 0 || err != nil {
			break
		}
		found_number, present := phone_book[query]
		if present == true {
			fmt.Printf("%s=%s\n", query, found_number)
		} else {
			fmt.Println("Not found")
		}
	}
}
