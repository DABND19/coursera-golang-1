package main

import "fmt"

func main() {
	var i int = 10
	var autoInt = -10

	fmt.Println(i, autoInt)

	var bigInt = 1<<32 - 1
	var unsignedInt uint = 100500
	var unsignedBigInt uint64 = 1<<64 - 1

	fmt.Println(bigInt, unsignedInt, unsignedBigInt)

	var f float32 = 1 + 2
	fmt.Println(f)

	var b bool
	var isOk bool = true
	var success = true

	fmt.Println(b, isOk, success)

	var c complex128 = -0.1 + 0.8i
	fmt.Println(c)
}
