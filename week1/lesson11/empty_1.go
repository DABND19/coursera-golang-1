package main

import (
	"fmt"
	"strconv"
)

type Payer interface {
	Pay(int) error
}

type Wallet struct {
	Cash int
}

func (w *Wallet) Pay(amount int) error {
	if w.Cash < amount {
		return fmt.Errorf("Not enough money in the wallet")
	}
	w.Cash -= amount
	return nil
}

func Buy(p Payer) {
	err := p.Pay(10)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Thank you for buy using %T\n\n", p)
}

func (w *Wallet) String() string {
	return "Wallet that contains " + strconv.Itoa(w.Cash) + " money."
}

func main() {
	myWallet := &Wallet{Cash: 100}
	fmt.Println(myWallet)
}
