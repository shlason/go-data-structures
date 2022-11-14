package sorting

import (
	"fmt"
	"strconv"
	"strings"
)

type list struct {
	items []int
}

func New() *list {
	return new(list)
}

func (l *list) Insert(e int) {
	l.items = append(l.items, e)
}

func (l *list) ToString() string {
	r := []string{}
	for _, item := range l.items {
		r = append(r, strconv.Itoa(item))
	}
	return strings.Join(r, ", ")
}

func (l *list) BubbleSort() {
	for i := 0; i < len(l.items); i++ {
		for j := 0; j < len(l.items)-1-i; j++ {
			if l.items[j] > l.items[j+1] {
				temp := l.items[j]
				l.items[j] = l.items[j+1]
				l.items[j+1] = temp
			}
		}
	}
}

func (l *list) SelectionSort() {
	for i := 0; i < len(l.items); i++ {
		indexOfMin := i
		for j := i; j < len(l.items); j++ {
			if l.items[j] < l.items[indexOfMin] {
				indexOfMin = j
			}
		}
		if indexOfMin != i {
			temp := l.items[i]
			l.items[i] = l.items[indexOfMin]
			l.items[indexOfMin] = temp
		}
	}
}

func (l *list) InsertionSort() {
	for i := 1; i < len(l.items); i++ {
		temp := l.items[i]
		j := i
		for j > 0 && l.items[j-1] > temp {
			l.items[j] = l.items[j-1]
			j--
		}
		l.items[j] = temp
	}
}

func (l *list) MergeSort() {
	l.items = mergeSortRec(l.items)
}

func mergeSortRec(list []int) []int {
	if len(list) < 2 {
		return list
	}
	mid := len(list) / 2
	left := make([]int, len(list[0:mid]))
	right := make([]int, len(list[mid:]))
	copy(left, list[0:mid])
	copy(right, list[mid:])
	fmt.Println(left, right)
	return merge(mergeSortRec(left), mergeSortRec(right))
}

func merge(left []int, right []int) []int {
	result := []int{}
	il := 0
	ir := 0

	for il < len(left) && ir < len(right) {
		if left[il] < right[ir] {
			result = append(result, left[il])
			il++
		} else {
			result = append(result, right[ir])
			ir++
		}
	}

	if il < len(left) {
		result = append(result, left[il:]...)
	}
	if ir < len(right) {
		result = append(result, right[ir:]...)
	}

	return result
}
