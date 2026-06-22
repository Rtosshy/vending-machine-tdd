package main

func main() {
}

type VendingMachine struct {
	buttons map[string]Button
	balance int
}

func (v *VendingMachine) InsertCoin(coin int) {
	v.balance += coin
}

func (v *VendingMachine) GetBalance() int {
	return v.balance
}

func (v *VendingMachine) PushButton(name string) error {
	return nil
}

func NewVendingMachine(buttons map[string]Button, initialBalance int) *VendingMachine {
	return &VendingMachine{buttons: buttons, balance: initialBalance}
}

type Button struct {
	name string
}

func (b *Button) Push() string {
	return b.name
}

func NewButton(name string) *Button {
	return &Button{name: name}
}
