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

func (l IntList) Set(i int, v Item) {
	l[i] = v.(int)
}

func (l IntList) Equal(target Lister) bool {
	if l.Len() != target.Len() {
		return false
	}
	r := true 
	Each(l, func(v Item, i int){
		if v != target.Get(i) {
			r = false
		}
	})
	return r 
}

func TestEach(t *testing.T) {
	arr := IntList([]int{1,3,5,7,9})
	Each(arr, func(v Item, i int){
		fmt.Printf("item %v has type of %T\n", v, v)
	})
}

func TestMap(t *testing.T) {
	arr := IntList([]int{1,3,5,7,9})
	newArr := Map(arr, func(v Item, i int) Item {
		return v.(int) * 2
	})
	if !IntList([]int{2, 6, 10, 14, 18}).Equal(newArr) {
		t.Errorf("before map %v \n after map %v", arr, newArr)
	}
}

func TestList(t *testing.T) {
	list := List{[]Item{1,2,3}}
	list2 := Map(list, func(v Item, i int) Item {
		return v.(int) * 2
	})
	
}