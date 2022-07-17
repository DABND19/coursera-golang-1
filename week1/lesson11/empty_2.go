package main

import "fmt"

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

func Buy(in interface{}) {
	p, ok := in.(Payer)
	if !ok {
		fmt.Printf("%T is not a payer\n\n", in)
		return
	}

	err := p.Pay(10)
	if err != nil {
		fmt.Printf("Payment error")
		return
	}
	fmt.Printf("Thank you for buy using %T\n\n", p)
}

func main() {
	myWallet := &Wallet{Cash: 100}
	Buy(myWallet)
	Buy([]int{1, 2, 3})
	Buy(3.1415)
}
