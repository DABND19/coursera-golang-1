package main

import "fmt"

func main() {
	a := 2
	b := &a
	*b = 3

	fmt.Println(a)

	c := &a

	d := new(int)
	*d = 12
	*c = *d
	*d = 13

	fmt.Println(a, *c, *d)

	c = d
	*c = 14
	fmt.Println(a, *c, *d)
}
