package sorting

import (
	"fmt"
	"testing"
)

func TestSorting(t *testing.T) {
	t.Error()

	list := createNonSortedList(9)
	fmt.Println(list)
	fmt.Println(list.BinarySearch(10))
	list.QuickSortV2()
	fmt.Println(list)
}

func createNonSortedList(size int) *list {
	r := New()
	for i := size; i > 0; i-- {
		r.Insert(i)
	}
	return r
}
