package main

import "fmt"

type Person struct {
	Id      int
	Name    string
	Address string
}

type Account struct {
	Id      int
	Name    string
	Cleaner func(string) string
	Owner   Person
	Person
}

func main() {
	acc := Account{
		Id:     1,
		Name:   "DABND",
		Person: Person{2, "David", "Moscow"},
	}
	fmt.Println(acc)
	fmt.Println(acc.Address)
	fmt.Println(acc.Name)
	fmt.Println(acc.Person.Name)

	acc.Owner = Person{2, "David Grigorenko", "Moscow, Russia"}
	fmt.Println(acc)
}
