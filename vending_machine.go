package main

import "errors"

type VendingMachine struct {
	buttons map[string]*Button
	balance int
}

func (v *VendingMachine) InsertCoin(coin int) {
	v.balance += coin
}

func (v *VendingMachine) GetBalance() int {
	return v.balance
}

func (v *VendingMachine) SelectButton(name string) (*Button, error) {
	button, ok := v.buttons[name]
	if !ok {
		return nil, errors.New("invalid button selected")
	}

	return button, nil
}

func (v *VendingMachine) Payment(price int) error {
	v.balance -= price
	return nil
}

func (v *VendingMachine) PushButton(name string) error {
	return nil
}

func NewVendingMachine(buttons map[string]*Button, initialBalance int) *VendingMachine {
	return &VendingMachine{buttons: buttons, balance: initialBalance}
}
