package list_test

import (
	"fmt"
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

func (l IntList) New() Lister {
	n := len(l)
	return IntList(make([]int, n))
}

func (l IntList) Set(i int, v Item) {
	l[i] = v.(int)
}

func TestEach(t *testing.T) {
	arr := IntList([]int{1,3,5,7,9})
	Each(arr, func(v Item, i int, l Lister){
		fmt.Printf("item %v has type of %T\n", v, v)
	})
}