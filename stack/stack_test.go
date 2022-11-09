package stack

import "testing"

func TestNewStack(t *testing.T) {
	s := New()

	if s.Size() != 0 {
		t.Errorf("Expected stack size is 0, but got %v", s.Size())
	}

	if !s.IsEmpty() {
		t.Errorf("Expected stack is empty, but got %v", s.Size())
	}

	datas := []int{1, 2, 3, 4, 5}

	for _, data := range datas {
		s.Push(data)
	}

	if s.Size() != len(datas) {
		t.Errorf("Expected stack size is %v, but got %v", len(datas), s.Size())
	}

	if s.IsEmpty() {
		t.Errorf("Expected stack is not empty")
	}

	if s.Peek() != datas[len(datas)-1] {
		t.Errorf("Expected stack peek get %v, but got %v", datas[len(datas)-1], s.Peek())
	}

	p1 := s.Pop()

	if p1 != datas[len(datas)-1] {
		t.Errorf("Expected stack pop get %v, but got %v", datas[len(datas)-1], p1)
	}

	p2 := s.Pop()

	if p2 != datas[len(datas)-2] {
		t.Errorf("Expected stack pop get %v, but got %v", datas[len(datas)-2], p2)
	}

	if s.Size() != 3 {
		t.Errorf("Expected stack size is %v, but got %v", 3, s.Size())
	}

	s.Clear()

	if s.Size() != 0 {
		t.Errorf("Expected stack size is 0, but got %v", s.Size())
	}

	if !s.IsEmpty() {
		t.Errorf("Expected stack is empty")
	}
}
