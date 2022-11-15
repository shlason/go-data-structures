package sorting

import (
	"fmt"
	"testing"
)

func TestSorting(t *testing.T) {
	t.Error()

	list := createNonSortedList(9)
	fmt.Println(list)
	list.QuickSort()
	fmt.Println(list)
}

func createNonSortedList(size int) *list {
	r := New()
	for i := size; i > 0; i-- {
		r.Insert(i)
	}
	return r
}
