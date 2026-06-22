package main

import (
	"testing"
)

func TestGetBalance(t *testing.T) {
	mockButtons := NewMockButtons()
	balance := 100
	v := NewVendingMachine(mockButtons, balance)
	got := v.GetBalance()
	expected := 100
	if got != expected {
		t.Errorf("got: %v, expected: %v", got, expected)
	}
}

func TestSelectButton(t *testing.T) {
}

func TestPayment(t *testing.T) {

}

func TestPushButton(t *testing.T) {

}
