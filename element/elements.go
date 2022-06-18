package element

type Elements[T Element] []T

func (es Elements[Element]) Len() int {
	return len(es)
}
