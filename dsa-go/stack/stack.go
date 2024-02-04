package stack

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(i T) {
	s.items = append(s.items, i)
}

func (s *Stack[T]) Pop() T {
	l := len(s.items) - 1
	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove
}

func (s *Stack[T]) Peek() T {
	return s.items[len(s.items)-1]
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack[T]) Size() int {
	return len(s.items)
}

func New[T any]() *Stack[T] {
	return &Stack[T]{}
}
