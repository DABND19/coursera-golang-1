package main

import "fmt"

type Person struct {
	Id   int
	Name string
}

func (p *Person) SetName(name string) {
	p.Name = name
}

type Account struct {
	Id   int
	Name string
	Person
}

func (acc *Account) SetName(name string) {
	acc.Name = name
}

type MySlice []int

func (sl *MySlice) GetLength() int {
	return len(*sl)
}

func main() {
	person := Person{1, "David"}
	person.SetName("David Grigorenko")
	fmt.Println(person)

	account := Account{
		Id:     1,
		Name:   "DABND",
		Person: Person{1, "David Grigorenko"},
	}
	account.SetName("grigorenko.david")
	account.Person.SetName("David")
	fmt.Println(account)

	mySlice := MySlice{1, 2, 3, 4, 5}
	fmt.Println(mySlice.GetLength())
}
