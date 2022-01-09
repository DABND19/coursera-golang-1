package main

import "fmt"

func deferTest() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Panic happend FIRST:", err)
		}
	}()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Panic happend SECOND:", err)
			panic("Second panic")
		}
	}()

	fmt.Println("Some useful work")
	panic("Something bad happends")
}

func main() {
	deferTest()
	return
}
