package main

import "fmt"

func main() {
	mapping := map[string]string{
		"key":  "value",
		"key1": "value1",
	}
	fmt.Println(mapping, len(mapping))

	ints := map[string]int{
		"key": 123,
	}
	fmt.Println(ints["key"], ints["not_existed_key"])

	_, isExists := ints["key"]
	fmt.Println(isExists)

	value, isExists := ints["not_existed_key"]
	fmt.Println(value, isExists)

	delete(mapping, "key1")
	fmt.Println(mapping)
}
