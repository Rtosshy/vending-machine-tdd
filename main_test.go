package main

import (
	"testing"
)

func TestPushButton(t *testing.T) {
	name := "cola"
	b := NewButton(name)
	got := b.Push()
	expected := "cola"

	if got != expected {
		t.Errorf("got: %v, expected: %v", got, expected)
	}
}

func TestInsertCoin(t *testing.T) {
	mock := NewMockButton()
	initialBalance := 0
	v := NewVendingMachine(*mock, initialBalance)
	input := 100
	v.InsertCoin(input)
	got := v.GetBalance()
	expected := 100
	if got != expected {
		t.Errorf("got: %v, expected: %v", got, expected)
	}

}

func TestGetBalance(t *testing.T) {
	mock := NewMockButton()
	balance := 100
	v := NewVendingMachine(*mock, balance)
	got := v.GetBalance()
	expected := 100
	if got != expected {
		t.Errorf("got: %v, expected: %v", got, expected)
	}
}
