package list_test

import (
	"testing"
	. "../list"
)


func TestEach(t *testing.T) {
	list := List{[]Item{1,3,5,7,9}}
	list2 := list.New(0)
	Each(&list, func(v Item, i int){
		list2.Append(v)
	})
	if !Equal(&list, list2) {
		t.Errorf("seems Each not works")
	}
}

func TestMap(t *testing.T) {
	list := List{[]Item{1,3,5,7,9}}
	newList := Map(&list, func(v Item, i int) Item {
		return v.(int) * 2
	})
	if !Equal(newList, &List{[]Item{2,6,10,14,18}}) {
		t.Errorf("before map %v \n after map %v", list, newList)
	}
}

func TestFilter(t *testing.T) {
	list := List{[]Item{1,2,3,4,5,6,7}}
	list2 := Filter(&list, func(v Item, i int) bool {
		return v.(int) % 2 == 0
	})
	if !Equal(list2, &List{[]Item{2,4,6}}){
		t.Errorf("Filter seems not work")
	}
}

func TestFindIndex(t *testing.T) {
	list := List{[]Item{'a', 'b', 'c', 'd', 'e'}}
	index := FindIndex(&list, func(v Item, i int) bool {
		return v.(rune) == 'e'
	})
	if index != 4 {
		t.Errorf("expect item index is 4 but got %v", index)
	}
}

func TestEqual(t *testing.T) {
	a := List{[]Item{'a', 'b', 'c', 'd', 'e'}}
	b := List{[]Item{'a', 'b', 'c', 'd', 'e'}}
	
	if !Equal(&a, &b) {
		t.Errorf("s %v should be equal to t %v, but got not equal", a, b)
	}
}