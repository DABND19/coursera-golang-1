package main

import "fmt"

func main() {
	boolValue := true
	if boolValue {
		fmt.Println("booValue is true")
	}

	user := map[string]string{
		"name": "DABND",
	}
	if _, isExists := user["name"]; isExists {
		fmt.Println("key \"name\" exists")
	}

	cond := 1
	if cond == 1 {
		fmt.Println("cond is 1")
	} else if cond == 2 {
		fmt.Println("cond is 2")
	}

	stringValue := "value"
	switch stringValue {
	case "value":
		fallthrough
	case "test", "value1":
		fmt.Println("hello, world!!!")
	default:
		fmt.Println("default case")
	}

	val1, val2 := 2, 2
	switch {
	case val1 > 1 || val2 < 3:
		fmt.Println("1st case")
	case val1 > 0:
		fmt.Println("2nd case")
	}

Loop:
	for key, val := range user {
		fmt.Println(key, val)
		switch {
		case key == "lastName":
			break
		case key == "name" && val == "DABND":
			break Loop
		}
	}
}
