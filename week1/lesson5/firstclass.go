package main

import "fmt"

func doNothing() {
	fmt.Println("I'm regular function")
}

func main() {
	func(in string) {
		fmt.Println("anon func out:", in)
	}("nobody")

	printer := func(in string) {
		fmt.Println("printer outs:", in)
	}
	printer("as variable")

	type strFuncType func(string)
	worker := func(callback strFuncType) {
		callback("as callback")
	}
	worker(printer)

	prefixer := func(prefix string) strFuncType {
		return func(in string) {
			fmt.Printf("[%s] %s\n", prefix, in)
		}
	}

	logLevel := "SUCCESS"
	successLoger := prefixer(logLevel)
	logLevel = "ERROR"
	successLoger("expected behaviour")
}
