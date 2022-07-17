package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var str string
	var helloWorld = "Hello, world!!!"
	var russianHelloWorld = "Привет, мир!!!"
	fmt.Println(str, helloWorld, russianHelloWorld)

	var rawBytes byte = '\x27'
	var literal rune = 'a'
	fmt.Println(rawBytes, literal)

	andGoodEvening := helloWorld + " And good evening."
	fmt.Println(andGoodEvening)

	byteLen := len(russianHelloWorld)
	symbolsCount := utf8.RuneCountInString(russianHelloWorld)
	fmt.Println(byteLen, symbolsCount)

	hello := russianHelloWorld[:12] // Привет - 0-11 байты!
	firstLiter := russianHelloWorld[0] // Нулевой байт, а не символ!
	fmt.Println(hello, firstLiter)

	byteString := []byte(helloWorld)
	helloWorld = string(byteString)
	fmt.Println(byteString, helloWorld)
}
