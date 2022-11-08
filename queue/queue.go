package queue

type element interface{}

type queue struct {
	items []element
}

func New() *queue {
	q := new(queue)
	q.items = make([]element, 0, 8)
	return q
}

func (q *queue) Enqueue(e element) {
	q.items = append(q.items, e)
}

func (q *queue) Dequeue() element {
	r := q.items[0]
	q.items = q.items[1:]
	return r
}

func (q *queue) Front() element {
	return q.items[0]
}

func (q *queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *queue) Size() int {
	return len(q.items)
}
