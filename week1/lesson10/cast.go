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

type Card struct {
	Balance    int
	ValidUntil string
	Cardholder string
	CVV        string
	Number     string
}

func (c *Card) Pay(amount int) error {
	if c.Balance < amount {
		return fmt.Errorf("Not enough money in the card")
	}
	c.Balance -= amount
	return nil
}

type ApplePay struct {
	Money   int
	AppleID string
}

func (a *ApplePay) Pay(amount int) error {
	if a.Money < amount {
		return fmt.Errorf("Not enough money in the account")
	}
	a.Money -= amount
	return nil
}

func Buy(p Payer) {
	switch p.(type) {
	case *Wallet:
		fmt.Println("Cash payment")
	case *Card:
		plasticCard, ok := p.(*Card)
		if !ok {
			fmt.Println("Can't cast to type \"Card\"")
			return
		}
		fmt.Println("Please, insert your card,", plasticCard.Cardholder)
	default:
		fmt.Println("Wow, something new!!!")
	}

	err := p.Pay(10)
	if err != nil {
		fmt.Printf("Error when paying by %T\n\n", p)
		return
	}
	fmt.Printf("Thank you for buy using %T\n\n", p)
}

func main() {
	myWallet := &Wallet{Cash: 100}
	Buy(myWallet)

	var myMoney Payer
	myMoney = &Card{Balance: 100, Cardholder: "DABND"}
	Buy(myMoney)

	myMoney = &ApplePay{Money: 9}
	Buy(myMoney)
}
