package main

import "fmt"

const pi = 3.1415

const (
	hello = "Привет"
	e = 2.7
)

const (
	zero = iota
	_
	two
)

const (
	_ = iota
	KB uint64 = 1 << (10 * iota)
	MB
)

const (
	year = 2021 // нетипизированная константа
	yearTyped int = 2021
)

func main() {
	fmt.Println(zero, two)

	fmt.Println(KB, MB)

	fmt.Println(year, yearTyped, year + yearTyped)

	var month int32 = 12
	fmt.Println(year + month)
	// fmt.Println(yearTyped + month) // (mismatched types int and int32)
}
