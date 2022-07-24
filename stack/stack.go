package stack

type Stack[T any] []T

func (s *Stack[T]) Push(v T) {
	*s = append(*s, v)
}

func (s *Stack[T]) Pop() *T {
	if len(*s) == 0 {
		return nil
	}

	v := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return &v
}

func NewStack[T any]() Stack[T] {
	return make(Stack[T], 0, 10)
}
