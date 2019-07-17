package list_test

import (
	"fmt"
	"testing"
	. "../list"
)


func TestEach(t *testing.T) {
	list := List{[]Item{1,3,5,7,9}}
	Each(&list, func(v Item, i int){
		fmt.Printf("item %v has type of %T\n", v, v)
	})
}

func TestMap(t *testing.T) {
	list := List{[]Item{1,3,5,7,9}}
	newList := Map(&list, func(v Item, i int) Item {
		return v.(int) * 2
	})
	
	t.Errorf("before map %v \n after map %v", list, newList)
	
}

func TestList(t *testing.T) {
	list := List{[]Item{1,2,3}}
	list2 := Map(&list, func(v Item, i int) Item {
		return v.(int) * 2
	})
	t.Errorf("%v %v", list2, list)
}

func TestFilter(t *testing.T) {
	list := List{[]Item{1,2,3,4,5,6,7}}
	list2 := Filter(&list, func(v Item, i int) bool {
		return v.(int) % 2 == 0
	})
	t.Errorf("%v", list2)
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