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
		return fmt.Errorf("нехватает денег")
	}
	w.Cash -= amount
	return nil
}

func main() {

	//// используем методы
	//wallet := Wallet{Cash: 100}
	//fmt.Println(wallet.Cash)
	//
	//wallet.Pay(25)
	//fmt.Println(wallet.Cash)

	// Используем через интерфейс
	myWallet := &Wallet{Cash: 100}
	Buy(myWallet)

}

func Buy(p Payer) {
	err := p.Pay(10)
	if err != nil {
		panic(err)
	}
	fmt.Println("спасибо за покупку")
}
