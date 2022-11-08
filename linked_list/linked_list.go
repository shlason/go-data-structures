package linkedList

import (
	"strings"
)

type element interface{}

type node struct {
	element
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func New() *linkedList {
	return new(linkedList)
}

func (ll *linkedList) Append(e element) {
	ptrN := &node{
		element: e,
	}
	ll.length++
	if ll.IsEmpty() {
		ll.head = ptrN
		return
	}
	current := ll.head
	for current.next != nil {
		current = current.next
	}
	current.next = ptrN
}

func (ll *linkedList) RemoveAt(index int) element {
	if !ll.isValidIndex(index) {
		return nil
	}
	var prev *node

	i := 0
	current := ll.head

	if index == 0 {
		ll.length--
		ll.head = current.next
		return current.element
	}

	for i != index {
		i++
		prev = current
		current = current.next
	}

	prev.next = current.next
	ll.length--

	return current.element
}

func (ll *linkedList) Insert(e element, index int) bool {
	if !ll.isValidIndex(index) {
		return false
	}
	var prev *node

	ptrN := &node{
		element: e,
	}
	i := 0
	current := ll.head

	if index == 0 {
		ll.head = ptrN
		ptrN.next = current
		ll.length++
		return true
	}

	for index != i {
		i++
		prev = current
		current = current.next
	}

	ptrN.next = current
	prev.next = ptrN
	ll.length++

	return true
}

func (ll *linkedList) IndexOf(e element) int {
	i := 0
	current := ll.head

	for current != nil {
		if current.element == e {
			return i
		}
		current = current.next
	}

	return -1
}

func (ll *linkedList) Remove(e element) bool {
	var prev *node
	i := 0
	current := ll.head

	for current != nil {
		if current.element == e {
			if i == 0 {
				ll.head = current.next
			} else {
				prev.next = current.next
			}
			return true
		}
		i++
		prev = current
		current = current.next
	}

	return false
}

func (ll *linkedList) ToString() string {
	var resultStrs []string

	current := ll.head

	for current != nil {
		resultStrs = append(resultStrs, current.element.(string))
		current = current.next
	}

	return strings.Join(resultStrs, " -> ")
}

func (ll *linkedList) Size() int {
	return ll.length
}

func (ll *linkedList) IsEmpty() bool {
	return ll.head == nil
}

func (ll *linkedList) isValidIndex(index int) bool {
	return index >= 0 && index < ll.length
}
