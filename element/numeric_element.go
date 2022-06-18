package element

type NumericElement float64

func (ne NumericElement) IsElement() bool {
	return true
}
