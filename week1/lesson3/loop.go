package main

import "fmt"

func main() {
	for {
		fmt.Println("loop iteration")
		break
	}

	isRunned := true
	for isRunned {
		fmt.Println("loop iteration with condition")
		isRunned = false
	}

	for i := 0; i < 2; i++ {
		if i == 0 {
			continue
		}
		fmt.Println(i, " loop iteration")
	}

	buf := []int{1, 2, 3}
	for index, item := range buf {
		fmt.Println(index, item)
	}

	mapping := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	for key, val := range mapping {
		fmt.Println(key, val)
	}

	str := "Привет, мир!!!"
	for pos, char := range str {
		fmt.Println(pos, char)
	}
}
