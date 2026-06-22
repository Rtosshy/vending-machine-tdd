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
	// coin := 100
}

func TestGetbalance(t *testing.T) {
	name := "test"
	mockButton := NewButton(name)
	balance := 100
	v := NewVendingMachine(*mockButton, balance)
	got := v.GetBalance()
	expected := 100
	if got != expected {
		t.Errorf("got: %v, expected: %v", got, expected)
	}
}
