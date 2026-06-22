package main

import (
	"errors"
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
	mockButton := NewMockButton()
	mock := NewButtons(mockButton)
	initialBalance := 0
	v := NewVendingMachine(mock, initialBalance)
	selectedDrink := "cola"

	_, err := v.SelectButton(selectedDrink)
	want := errors.New("invalid button selected")

	if err == nil {
		t.Errorf("got: %v, want: %v", err, want)
	}

	selectedDrink = "mock"
	_, err = v.SelectButton(selectedDrink)
	want = nil

	if err != nil {
		t.Errorf("got: %v, want: %v", err, want)
	}
}

func TestPayment(t *testing.T) {
}

func TestPushButton(t *testing.T) {
}
