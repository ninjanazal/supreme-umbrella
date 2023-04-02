package classes

// Queue is a generic data structure that implements a FIFO (first-in, first-out) queue.
// It is implemented as a slice of type T, where T can be any type.
type Queue[T any] []T

// NewQueue creates a new empty Queue instance and returns a pointer to it.
// @return {*Queue[T]}: A pointer to the new empty Queue instance.
func NewQueue[T any]() *Queue[T] {
	que := make(Queue[T], 0)
	return &que
}

// Add adds a new element of type T to the end of the queue.
// @param {T} v: The value to be added to the queue.
func (q *Queue[T]) Add(v T) {
	*q = append(*q, v)
}

// Next removes and returns the first element in the queue.
// @return {T}: The first element in the queue.
// If the queue is empty, a new zero value of type T will be returned.
func (q *Queue[T]) Next() T {
	if len(*q) == 0 {
		return *new(T)
	}
	var obj T = (*q)[0]
	if len(*q) == 1 {
		*q = make(Queue[T], 0)
	} else {
		*q = (*q)[1:len(*q)]
	}

	return obj
}

// PeekNext returns the first element in the queue without removing it.
// @return {T}: The first element in the queue.
// If the queue is empty, a new zero value of type T will be returned.
func (q *Queue[T]) PeekNext() T {
	if len(*q) == 0 {
		return *new(T)
	}
	return (*q)[0]
}

// Len returns the number of elements in the queue.
// @return {int}: The number of elements in the queue.
func (q *Queue[T]) Len() int {
	return len(*q)
}

// IsEmpty checks if the queue is empty.
// @return {bool}: Returns true if the queue is empty, false otherwise.
func (q *Queue[T]) IsEmpty() bool {
	return len(*q) == 0
}
