package main

func NewMockButton() *Button {
	return &Button{name: "mock"}
}

func NewMockButtons() map[string]*Button {
	mock := NewMockButton()
	buttons := make(map[string]*Button)
	buttons[mock.name] = mock
	return buttons
}
