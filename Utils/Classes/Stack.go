package classes

// Stack is a generic implementation of a stack data structure.
// It is a slice of elements of type T, where T can be any type.
// The zero value of a Stack is an empty stack.
type Stack[T any] []T

// NewStack creates a new empty stack and returns a pointer to it.
// @return {*Stack[T]} - A pointer to a new empty stack of type T.
func NewStack[T any]() *Stack[T] {
	st := make(Stack[T], 0)
	return &st
}

// Push adds an element to the top of the stack.
// @param {T} v - The value of type T to be added to the stack.
func (s *Stack[T]) Push(v T) {
	*s = append(*s, v)
}

// Pop removes and returns the top element of the stack.
// If the stack is empty, it returns the zero value of type T.
// @return {T} - The top element of the stack.
func (s *Stack[T]) Pop() T {
	if len(*s) == 0 {
		return *new(T)
	}
	var obj T = (*s)[len(*s)-1]
	*s = (*s)[:len(*s)]
	return obj
}

// Peek returns the top element of the stack without removing it.
// If the stack is empty, it returns the zero value of type T.
// @return {T} - The top element of the stack.
func (s *Stack[T]) Peek() T {
	if len(*s) == 0 {
		return *new(T)
	}
	return (*s)[len(*s)-1]
}

// Len returns the number of elements in the stack.
// @return {int} - The number of elements in the stack.
func (s *Stack[T]) Len() int {
	return len(*s)
}

// IsEmpty returns true if the stack is empty, false otherwise.
// @return {bool} - True if the stack is empty, false otherwise.
func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}
