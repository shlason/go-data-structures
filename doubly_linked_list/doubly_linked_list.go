package doublyLinkedList

import "strings"

type element interface{}

type node struct {
	element
	prev *node
	next *node
}

type doublyLinkedList struct {
	length int
	head   *node
	tail   *node
}

func New() *doublyLinkedList {
	return new(doublyLinkedList)
}

func (dll *doublyLinkedList) Insert(e element, index int) bool {
	if index < 0 || index > dll.Size() {
		return false
	}
	var prev *node
	i := 0
	current := dll.head
	ptrN := &node{
		element: e,
	}

	if index == 0 {
		if dll.IsEmpty() {
			dll.head = ptrN
			dll.tail = ptrN
		} else {
			ptrN.next = current
			current.prev = ptrN
			dll.head = ptrN
		}
		dll.length++
		return true
	}

	if index == dll.Size() {
		current = dll.tail
		ptrN.prev = current
		current.next = ptrN
		dll.tail = ptrN
		dll.length++
		return true
	}

	for i != index {
		i++
		prev = current
		current = current.next
	}

	ptrN.next = current
	ptrN.prev = prev
	current.prev = ptrN
	prev.next = ptrN

	dll.length++
	return true
}

func (dll *doublyLinkedList) RemoveAt(index int) element {
	if index < 0 || index >= dll.Size() {
		return nil
	}

	var prev *node
	i := 0
	current := dll.head

	if index == 0 {
		dll.head = current.next
		if dll.head == nil {
			dll.tail = nil
		} else {
			dll.head.prev = nil
		}
		dll.length--
		return current.element
	}

	if index == dll.Size()-1 {
		current = dll.tail
		dll.tail = current.prev
		dll.tail.next = nil
		dll.length--
		return current.element
	}

	for i != index {
		i++
		prev = current
		current = current.next
	}

	prev.next = current.next
	current.next.prev = prev
	dll.length--
	return current.element
}

func (dll *doublyLinkedList) ToString() string {
	var resultStrs []string

	current := dll.head

	for current != nil {
		resultStrs = append(resultStrs, current.element.(string))
		current = current.next
	}

	return strings.Join(resultStrs, " <-> ")
}

func (dll *doublyLinkedList) Size() int {
	return dll.length
}

func (dll *doublyLinkedList) IsEmpty() bool {
	return dll.Size() == 0
}
