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
	if !reflect.DeepEqual(list2, expectedList) {
		t.Errorf("intlist Map error: list2 should be %v", expectedList)
	}
}

func TestIntListFilter(t *testing.T) {
	list1 := IntList([]int{0,1,2,3,4,5,6})
	list2 := list1.Filter(func(v, _ int) bool {
		return v % 2 == 0
	})
	expectedList := IntList([]int{0,2,4,6})
	if !reflect.DeepEqual(list2, expectedList) {
		t.Errorf("intlist Map error: list2 should be %v", expectedList)
	}
}

func TestIntListFindIndex(t *testing.T) {
	list := IntList([]int{5,8,13,17})
	index := list.FindIndex(17)
	if index != 3 {
		t.Error("intlist FindIndex error: the index should be 3")
	}
}

func TestIntListContains(t *testing.T) {
	list := IntList([]int{5,8,13,17})
	if !list.Contains(13) {
		t.Error("intlist Contains error: list should contains 13")
	}
}

func TestIntListReduce(t *testing.T) {
	list := IntList([]int{5,8,13,17})
	sum := list.Reduce(func(a, b int) int {
		return a + b
	}, 0)
	if sum != 43 {
		t.Error("intlist Reduce error: sum should be 43")
	}
}