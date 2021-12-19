package main

import "fmt"

func main() {
	var num0 int
	fmt.Println(num0)

	var num1 int = 1
	fmt.Println(num1)

	var num2 = 2
	fmt.Println(num2)

	num3 := 3
	fmt.Println(num3)

	num3 += 1
	num3++
	fmt.Println(num3)

	weight, height := 20, 30
	fmt.Println(weight, height)
}
