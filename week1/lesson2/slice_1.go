package main

import "fmt"

func main() {
	var buf0 []int
	buf1 := []int{}
	buf2 := []int{42}
	buf3 := make([]int, 0)
	buf4 := make([]int, 5)
	buf5 := make([]int, 5, 10)
	fmt.Println(buf0, buf1, buf2, buf3, buf4, buf5)

	var buf []int
	buf = append(buf, 9, 10)
	buf = append(buf, 12)
	fmt.Println(buf)

	otherBuf := make([]int, 3)
	buf = append(buf, otherBuf...)
	fmt.Println(buf, otherBuf)
	fmt.Println(len(buf), cap(buf))
}
