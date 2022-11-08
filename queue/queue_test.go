package queue

import "testing"

func TestQueue(t *testing.T) {
	q := New()

	if q.Size() != 0 {
		t.Errorf("Expected queue size is 0, but got %v", q.Size())
	}

	if !q.IsEmpty() {
		t.Errorf("Expected queue is empty")
	}

	datas := []string{"testing1", "testing2", "testing3"}

	for _, data := range datas {
		q.Enqueue(data)
	}

	if q.Size() != len(datas) {
		t.Errorf("Expected queue size is %v, but got %v", len(datas), q.Size())
	}

	if q.IsEmpty() {
		t.Errorf("Expected queue is not empty")
	}

	d1 := q.Dequeue()

	if d1 != datas[0] {
		t.Errorf("Expected dequeue element is %v, but got %v", datas[0], d1)
	}

	f1 := q.Front()

	if f1 != datas[1] {
		t.Errorf("Expected front element is %v, but got %v", datas[1], f1)
	}

	if q.Size() != 2 {
		t.Errorf("Expected queue size is %v, but got %v", 1, q.Size())
	}
}
