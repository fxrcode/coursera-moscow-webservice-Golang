package main

import (
	"fmt"
)

type Payer interface {
	Pay(int) error
}

type Wallet struct {
	Cash int
}

func (w *Wallet) Pay(amount int) error {
	if w.Cash < amount {
		return fmt.Errorf("Не хватает денег в кошельке")
	}
	w.Cash -= amount
	return nil
}

func Buy(p Payer, price int) {
	fmt.Printf("Detail %#v\n\n", p)

	err := p.Pay(price)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Спасибо за покупку через %T, detail %#v\n\n", p, p)
}

func main() {
	myWallet := &Wallet{Cash: 100}
	Buy(myWallet, 10)
}
