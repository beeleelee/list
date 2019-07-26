package list_test

import (
	"testing"
	"reflect"
	. "github.com/beeleelee/list"
)

func TestF64ListEach(t *testing.T) {
	list := F64List([]float64{2,3,5,7,9,11,13})
	list.Each(func(v float64, i int){
		t.Log(v, i)
	})
}

func TestF64ListMap(t *testing.T) {
	list1 := F64List([]float64{2,3,5,7,9,11,13})
	list2 := list1.Map(func(v float64, _ int) float64 {
		return v + 1
	})
	expectedList := F64List([]float64{3,4,6,8,10,12,14})
	if !reflect.DeepEqual(list2, expectedList) {
		t.Errorf("float64list Map error: list2 should be %v", expectedList)
	}
}

func TestF64ListFilter(t *testing.T) {
	list1 := F64List([]float64{0,1,2,3,4,5,6})
	list2 := list1.Filter(func(v float64, _ int) bool {
		return int(v) % 2 == 0
	})
	expectedList := F64List([]float64{0,2,4,6})
	if !reflect.DeepEqual(list2, expectedList) {
		t.Errorf("float64list Map error: list2 should be %v", expectedList)
	}
}

func TestF64ListFindIndex(t *testing.T) {
	list := F64List([]float64{5,8,13,17})
	index := list.FindIndex(17)
	if index != 3 {
		t.Error("float64list FindIndex error: the index should be 3")
	}
}

func TestF64ListContains(t *testing.T) {
	list := F64List([]float64{5,8,13,17})
	if !list.Contains(13) {
		t.Error("float64list Contains error: list should contains 13")
	}
}

func TestF64ListReduce(t *testing.T) {
	list := F64List([]float64{5,8,13,17})
	sum := list.Reduce(func(a, b float64) float64 {
		return a + b
	}, 0)
	if sum != 43 {
		t.Error("float64list Reduce error: sum should be 43")
	}
}

func TestF64ListSome(t *testing.T) {
	list := F64List([]float64{1,2,3,4,5})
	if !list.Some(func(v float64, i int) bool {
		return int(v) % 2 == 0
	}) {
		t.Error("float64list Some eror: list should have even value")
	}
}

func TestF64ListEvery(t *testing.T) {
	list := F64List([]float64{1,2,3,4,5})
	if list.Every(func(v float64, i int) bool {
		return int(v) % 2 == 0
	}) {
		t.Error("float64list Every eror: list should have odd value")
	}
}

func TestF64ListShuffle(t *testing.T) {
	list := F64List([]float64{1,2,3,4,5,6,7,8,9})
	t.Log(list.Shuffle())
}

func TestF64ListSort(t *testing.T) {
	list := F64List([]float64{3,9,2,5,1,7})
	list.Sort()
	t.Log(list)
}

func TestF64ListIsSorted(t *testing.T) {
	list := F64List([]float64{3,9,2,5,1,7})
	listSorted := F64List([]float64{1,2,3,4,5,6,7,8,9})
	if list.IsSorted() {
		t.Error("float64list IsSorted error: list should not have been sorted")
	}
	if !listSorted.IsSorted() {
		t.Error("float64list IsSorted error: listSorted should have been sorted")
	}
}

func TestF64ListGet(t *testing.T) {
	list := F64List([]float64{0,1,2,3,4,5,6,7,8,9})
	if item, _ := list.Get(-1); item != 9 {
		t.Error("float64list Get error, item should be 9")
	}
	if _, ok := list.Get(20); ok != false {
		t.Error("float64list Get error, ok should be false")
	}
	if item, _ := list.Get(-7); item != 3 {
		t.Error("float64list Get error, item should be 3")
	}
	if _, ok := list.Get(-20); ok != false {
		t.Error("float64list Get error, ok should be false")
	}
	
}

func TestF64ListTail(t *testing.T) {
	list := F64List([]float64{0,1,2,3,4,5,6,7,8,9})
	if !reflect.DeepEqual(list.Tail(1), F64List([]float64{9})) {
		t.Error("float64list Tail error, it should be [9]")
	}
	if !reflect.DeepEqual(list.Tail(3), F64List([]float64{7,8,9})) {
		t.Error("float64list Tail error, it should be [7,8,9]")
	}
}

func TestF64ListUnion(t *testing.T) {
	list1 := F64List([]float64{1,2,3})
	list2 := F64List([]float64{4,5,6})
	list3 := F64List([]float64{1,2,3,4,5,6})
	if !reflect.DeepEqual(list3, list1.Union(list2)) {
		t.Error("float64list Union error, list1 union list2 should be equal to list3")
	}
}

func TestF64ListIntersection(t *testing.T) {
	l1 := F64List([]float64{0,1,2,4,7})
	l2 := F64List([]float64{3,4,5,6,9,8})
	l3 := F64List([]float64{1,6,5,0})
	t.Log(l2.Intersection(l1))
	t.Log(l2.Intersection(l3))
	t.Log(l3.Intersection(l1))
	if !reflect.DeepEqual(l1.Intersection(l2), F64List([]float64{4})) {
		t.Error("float64list Intersection error, intersect list should be [4]")
	}
	if !reflect.DeepEqual(l2.Intersection(l3), F64List([]float64{5,6})) {
		t.Error("float64list Intersection error, intersect list should be [5,6]")
	}
	if !reflect.DeepEqual(l3.Intersection(l1), F64List([]float64{1,0})) {
		t.Error("float64list Intersection error, intersect list should be [1,0]")
	}
}

func TestF64ListDifference(t *testing.T) {
	l1 := F64List([]float64{0,1,2,4,7})
	l2 := F64List([]float64{3,4,5,6,9,8})
	l3 := F64List([]float64{1,6,5,0})
	t.Log(l2.Difference(l1))
	t.Log(l2.Difference(l3))
	t.Log(l3.Difference(l1))
	if !reflect.DeepEqual(l1.Difference(l2), F64List([]float64{0,1,2,7})) {
		t.Error("float64list Deference error, difference list should be [4]")
	}
	if !reflect.DeepEqual(l2.Difference(l3), F64List([]float64{3,4,9,8})) {
		t.Error("float64list Deference error, difference list should be [5,6]")
	}
	if !reflect.DeepEqual(l3.Difference(l1), F64List([]float64{6,5})) {
		t.Error("float64list Deference error, difference list should be [1,0]")
	}
}