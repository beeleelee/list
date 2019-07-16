package list_test

import (
	"testing"
	. "../list"
)

type IntList []int 

func (l IntList) Len() int {
	return len(l)
}

func (l IntList) Get(i int) Item {
	return l[i]
}

func (l IntList) New() Item {
	n := len(ic)
	return IntCn(make([]int, n))
}

func (l IntList) Set(i int, v Item) {
	l[i] = v.(int)
}

func TestEach(t *testing.T) {
	arr := IntCn([]int{1,3,5,7,9})
	Each(arr, func(v Item, i int, l Lister){
		t.Errorf("each item %v", v)
	})
}