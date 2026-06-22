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
