package element

type StringElement string

func (se StringElement) IsElement() bool {
	return true
}
