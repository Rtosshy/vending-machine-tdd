package main

func main() {

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
