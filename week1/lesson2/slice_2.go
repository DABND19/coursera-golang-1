package main

import "fmt"

func main() {
	buf := []int{1, 2, 3, 4, 5}
	fmt.Println(buf)

	fmt.Println(buf[1:4])
	fmt.Println(buf[:2])
	fmt.Println(buf[2:])

	newBuf := buf[:]
	fmt.Println(newBuf)
	newBuf[0] = 9
	fmt.Println(buf) // same data

	// reallocation of newBuf
	newBuf = append(newBuf, 6)
	fmt.Println(buf, newBuf) // different data

	copiedBuf := make([]int, len(buf), cap(buf))
	copy(copiedBuf, buf)
	fmt.Println(buf, copiedBuf)
	copiedBuf[0] = 123
	fmt.Println(buf, copiedBuf)

	ints := []int{1, 2, 3, 4}
	copy(ints[1:3], []int{5, 6})
	fmt.Println(ints)
}
