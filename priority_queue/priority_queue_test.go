package priorityQueue

import (
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	pq := NewPriority()

	if pq.Size() != 0 {
		t.Errorf("Expected priority queue size is 0, but got %v", pq.Size())
	}

	if !pq.IsEmpty() {
		t.Errorf("Expected priority queue is empty")
	}

	datas := []string{"1", "22", "55555", "333"}

	for _, data := range datas {
		pq.Enqueue(data, len(data))
	}

	if pq.Size() != len(datas) {
		t.Errorf("Expected priority queue size is %v, but got %v", len(datas), pq.Size())
	}

	if pq.IsEmpty() {
		t.Errorf("Expected priority queue is not empty")
	}

	d1 := pq.Dequeue()

	if d1 != datas[0] {
		t.Errorf("Expected dequeue element is %v, but got %v", datas[0], d1)
	}

	pq.Dequeue()
	d2 := pq.Dequeue()

	if d2 != datas[3] {
		t.Errorf("Expected dequeue element is %v, but got %v", datas[3], d2)
	}

	if pq.Front() != datas[2] {
		t.Errorf("Expected front element is %v, but got %v", datas[2], pq.Front())
	}

	if pq.Size() != 1 {
		t.Errorf("Expected priority queue size is 1, but got %v", pq.Size())
	}

	if pq.IsEmpty() {
		t.Errorf("Expected priority queue is not empty")
	}
}
