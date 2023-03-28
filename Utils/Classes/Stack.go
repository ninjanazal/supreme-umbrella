package classes

type Stack[T any] []T

func NewStack[T any]() *Stack[T] {
	st := make(Stack[T], 0)
	return &st
}

func (s *Stack[T]) Push(v T) {
	*s = append(*s, v)
}

func (s *Stack[T]) Pop() T {
	if len(*s) == 0 {
		return *new(T)
	}
	

}
