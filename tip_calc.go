package main

import (
	"fmt"
)

func main() {

	var mealCost float64
	var tipPercent float64
	var taxPercent float64

	fmt.Scan(&mealCost)
	fmt.Scan(&tipPercent)
	fmt.Scan(&taxPercent)

	tax := mealCost * (taxPercent / 100)
	tip := mealCost * (tipPercent / 100)

	fmt.Printf("The total meal cost is %.0f dollars.\n", mealCost+tax+tip)
}
