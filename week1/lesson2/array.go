package main

import "fmt"

func main() {
	var a1 [3]int
	fmt.Println(a1)

	const size = 2
	var a2 [2 * size]bool
	fmt.Println(a2)

	a3 := [...]int{1, 2, 3}
	fmt.Println(a3)
}
