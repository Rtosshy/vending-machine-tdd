package main

func main() {
}

type VendingMachine struct {
	button  Button
	balance int
}

func (v *VendingMachine) InsertCoin(coin int) {
	v.balance += coin
}

func (v *VendingMachine) GetBalance() int {
	return v.balance
}

func NewVendingMachine(button Button, balance int) *VendingMachine {
	return &VendingMachine{button: button, balance: balance}
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
