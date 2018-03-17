package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	rand.Seed(42)
	fmt.Println("My favorite number is", rand.Intn(10))

	var i uint64 = 4
	var d float64 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)

	var j uint64
	var e float64
	var t string

	scanner.Scan()
	j, _ = strconv.ParseUint(scanner.Text(), 10, 64)
	scanner.Scan()
	e, _ = strconv.ParseFloat(scanner.Text(), 64)
	scanner.Scan()
	t = scanner.Text()

	fmt.Println(strconv.FormatUint(i+j, 16))
	fmt.Println(d + e)
	fmt.Println(s, t)

}
