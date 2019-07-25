package list_test

import (
	"testing"
	"reflect"
	. "../list"
)

func TestIntListEach(t *testing.T) {
	list := IntList([]int{2,3,5,7,9,11,13})
	list.Each(func(v int, i int){
		t.Log(v, i)
	})
}

func TestIntListMap(t *testing.T) {
	list1 := IntList([]int{2,3,5,7,9,11,13})
	list2 := list1.Map(func(v, _ int) int {
		return v + 1
	})
	expectedList := IntList([]int{3,4,6,8,10,12,14})
	if reflect.DeepEqual(list2, expectedList) {
		t.Errorf("intlist map error: list2 should be %v", expectedList)
	}
}