package main

type Button struct {
	name string
}

func (b *Button) Push() string {
	return b.name
}

func NewButton(name string) *Button {
	return &Button{name: name}
}

func NewButtons(buttons ...*Button) map[string]*Button {
	button_map := make(map[string]*Button)
	for _, b := range buttons {
		button_map[b.name] = b
	}
	return button_map
}
