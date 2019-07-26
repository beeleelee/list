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

func TestIntListSome(t *testing.T) {
	list := IntList([]int{1,2,3,4,5})
	if !list.Some(func(v, i int) bool {
		return v % 2 == 0
	}) {
		t.Error("intlist Some eror: list should have even int")
	}
}

func TestIntListEvery(t *testing.T) {
	list := IntList([]int{1,2,3,4,5})
	if list.Every(func(v, i int) bool {
		return v % 2 == 0
	}) {
		t.Error("intlist Every eror: list should have odd int")
	}
}

func TestIntListShuffle(t *testing.T) {
	list := IntList([]int{1,2,3,4,5,6,7,8,9})
	t.Log(list.Shuffle())
}

func TestIntListSort(t *testing.T) {
	list := IntList([]int{3,9,2,5,1,7})
	list.Sort()
	t.Log(list)
}

func TestIntListIsSorted(t *testing.T) {
	list := IntList([]int{3,9,2,5,1,7})
	listSorted := IntList([]int{1,2,3,4,5,6,7,8,9})
	if list.IsSorted() {
		t.Error("intlist IsSorted error: list should not have been sorted")
	}
	if !listSorted.IsSorted() {
		t.Error("intlist IsSorted error: listSorted should have been sorted")
	}
}

func TestIntListGet(t *testing.T) {
	list := IntList([]int{0,1,2,3,4,5,6,7,8,9})
	if item, _ := list.Get(-1); item != 9 {
		t.Error("intlist Get error, item should be 9")
	}
	if _, ok := list.Get(20); ok != false {
		t.Error("intlist Get error, ok should be false")
	}
	if item, _ := list.Get(-7); item != 3 {
		t.Error("intlist Get error, item should be 3")
	}
	if _, ok := list.Get(-20); ok != false {
		t.Error("intlist Get error, ok should be false")
	}
	
}

func TestIntListTail(t *testing.T) {
	list := IntList([]int{0,1,2,3,4,5,6,7,8,9})
	if !reflect.DeepEqual(list.Tail(1), IntList([]int{9})) {
		t.Error("intlist Tail error, it should be [9]")
	}
	if !reflect.DeepEqual(list.Tail(3), IntList([]int{7,8,9})) {
		t.Error("intlist Tail error, it should be [7,8,9]")
	}
}

func TestIntListUnion(t *testing.T) {
	list1 := IntList([]int{1,2,3})
	list2 := IntList([]int{4,5,6})
	list3 := IntList([]int{1,2,3,4,5,6})
	if !reflect.DeepEqual(list3, list1.Union(list2)) {
		t.Error("intlist Union error, list1 union list2 should be equal to list3")
	}
}

func TestIntListIntersection(t *testing.T) {
	l1 := IntList([]int{0,1,2,4,7})
	l2 := IntList([]int{3,4,5,6,9,8})
	l3 := IntList([]int{1,6,5,0})
	t.Log(l2.Intersection(l1))
	t.Log(l2.Intersection(l3))
	t.Log(l3.Intersection(l1))
	if !reflect.DeepEqual(l1.Intersection(l2), IntList([]int{4})) {
		t.Error("intlist Intersection error, intersect list should be [4]")
	}
	if !reflect.DeepEqual(l2.Intersection(l3), IntList([]int{5,6})) {
		t.Error("intlist Intersection error, intersect list should be [5,6]")
	}
	if !reflect.DeepEqual(l3.Intersection(l1), IntList([]int{1,0})) {
		t.Error("intlist Intersection error, intersect list should be [1,0]")
	}
}

func TestIntListDifference(t *testing.T) {
	l1 := IntList([]int{0,1,2,4,7})
	l2 := IntList([]int{3,4,5,6,9,8})
	l3 := IntList([]int{1,6,5,0})
	t.Log(l2.Difference(l1))
	t.Log(l2.Difference(l3))
	t.Log(l3.Difference(l1))
	if !reflect.DeepEqual(l1.Difference(l2), IntList([]int{0,1,2,7})) {
		t.Error("intlist Deference error, difference list should be [4]")
	}
	if !reflect.DeepEqual(l2.Difference(l3), IntList([]int{3,4,9,8})) {
		t.Error("intlist Deference error, difference list should be [5,6]")
	}
	if !reflect.DeepEqual(l3.Difference(l1), IntList([]int{6,5})) {
		t.Error("intlist Deference error, difference list should be [1,0]")
	}
}