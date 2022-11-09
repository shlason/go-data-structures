package priorityQueue

type element interface{}

type priorityElement struct {
	element
	priority int
}

type priorityQueue struct {
	items []priorityElement
}

func NewPriority() *priorityQueue {
	pq := new(priorityQueue)
	pq.items = make([]priorityElement, 0, 8)
	return pq
}

func (pq *priorityQueue) Enqueue(e element, p int) {
	if pq.IsEmpty() {
		pq.items = append(pq.items, priorityElement{
			element:  e,
			priority: p,
		})
		return
	}
	var index int
	for i, pe := range pq.items {
		index = i
		if pe.priority > p {
			break
		}
	}
	item := priorityElement{
		element:  e,
		priority: p,
	}
	if pq.items[index].priority > p {
		itemsA := pq.items[:index]
		itemsB := make([]priorityElement, len(pq.items[index:]))

		copy(itemsB, pq.items[index:])

		itemsA = append(itemsA, item)
		pq.items = append(itemsA, itemsB...)
		return
	}
	pq.items = append(pq.items, item)
}

func (pq *priorityQueue) Dequeue() element {
	r := pq.Front()
	pq.items = pq.items[1:]
	return r
}

func (pq *priorityQueue) Front() element {
	return pq.items[0].element
}

func (pq *priorityQueue) Size() int {
	return len(pq.items)
}

func (pq *priorityQueue) IsEmpty() bool {
	return pq.Size() == 0
}
