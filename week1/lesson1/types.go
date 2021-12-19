package main

import "fmt"

type UserId int

func main() {
	idx := 1
	var uid UserId = 42

	// uid = idx // cannot use idx (type int) as type UserId in assignment
	uid = UserId(idx)
	fmt.Println(uid, idx)
}
