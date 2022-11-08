package stack

type element interface{}

type stack struct {
	items []element
}

func New() *stack {
	s := new(stack)
	s.items = make([]element, 0, 8)
	return s
}

func (s *stack) Push(e element) {
	s.items = append(s.items, e)
}

func (s *stack) Pop() element {
	r := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return r
}

func (s *stack) Peek() element {
	return s.items[len(s.items)-1]
}

func (s *stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *stack) Size() int {
	return len(s.items)
}

func (s *stack) Clear() {
	s.items = []element{}
}
