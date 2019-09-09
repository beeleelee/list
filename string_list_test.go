package list_test

import (
	"reflect"
	"testing"

	. "github.com/beeleelee/list"
)

func TestStrListEach(t *testing.T) {
	list := StrList([]string{"alex", "leo", "lily"})
	list.Each(func(v string, i int) {
		t.Log(v, i)
	})
}

func TestStrListMap(t *testing.T) {
	list1 := StrList([]string{"alex", "leo", "lily"})
	list2 := list1.Map(func(v string, _ int) string {
		return "Hi, " + v
	})
	expectedList := StrList([]string{"Hi, alex", "Hi, leo", "Hi, lily"})
	if !reflect.DeepEqual(list2, expectedList) {
		t.Errorf("stringlist Map error: list2 should be %v", expectedList)
	}
}

func TestStrListFilter(t *testing.T) {
	list1 := StrList([]string{"apple", "banana", "orange"})
	list2 := list1.Filter(func(v string, _ int) bool {
		return v == "apple"
	})
	expectedList := StrList([]string{"apple"})
	if !reflect.DeepEqual(list2, expectedList) {
		t.Errorf("stringlist Map error: list2 should be %v", expectedList)
	}
}

func TestStrListFindIndex(t *testing.T) {
	list := StrList([]string{"alex", "leo", "lily"})
	index := list.FindIndex("lily")
	if index != 2 {
		t.Error("stringlist FindIndex error: the index should be 2")
	}
}

func TestStrListContains(t *testing.T) {
	list := StrList([]string{"alex", "leo", "lily"})
	if !list.Contains("alex") {
		t.Error("stringlist Contains error: list should contains alex")
	}
}

func TestStrListSome(t *testing.T) {
	list := StrList([]string{"alex", "leo", "lily"})
	if !list.Some(func(v string, i int) bool {
		return v == "alex"
	}) {
		t.Error("stringlist Some eror: list should have alex")
	}
}

func TestStrListEvery(t *testing.T) {
	list := StrList([]string{"alex", "leo", "lily"})
	if list.Every(func(v string, i int) bool {
		return v == "leo"
	}) {
		t.Error("stringlist Every eror: list should only have one leo")
	}
}

func TestStrListShuffle(t *testing.T) {
	list := StrList([]string{"alex", "leo", "lily"})
	t.Log(list.Shuffle())
}

func TestStrListSort(t *testing.T) {
	list := StrList([]string{"alex", "leo", "lily"})
	list.Sort()
	t.Log(list)
}

func TestStrListIsSorted(t *testing.T) {
	list := StrList([]string{"a", "f", "c", "g", "e"})
	listSorted := StrList([]string{"a", "b", "c", "d", "e"})
	if list.IsSorted() {
		t.Error("stringlist IsSorted error: list should not have been sorted")
	}
	if !listSorted.IsSorted() {
		t.Error("stringlist IsSorted error: listSorted should have been sorted")
	}
}

func TestStrListGet(t *testing.T) {
	list := StrList([]string{"a", "b", "c", "d", "e"})
	if item, _ := list.Get(-1); item != "e" {
		t.Error("stringlist Get error, item should be e")
	}
	if _, ok := list.Get(20); ok != false {
		t.Error("stringlist Get error, ok should be false")
	}
	if item, _ := list.Get(-3); item != "c" {
		t.Error("stringlist Get error, item should be c")
	}
	if _, ok := list.Get(-20); ok != false {
		t.Error("stringlist Get error, ok should be false")
	}

}

func TestStrListTail(t *testing.T) {
	list := StrList([]string{"a", "b", "c", "d", "e"})
	if !reflect.DeepEqual(list.Tail(1), StrList([]string{"e"})) {
		t.Error("stringlist Tail error, it should be [e]")
	}
	if !reflect.DeepEqual(list.Tail(3), StrList([]string{"c", "d", "e"})) {
		t.Error("stringlist Tail error, it should be [c,d,e]")
	}
}

func TestStrListUnion(t *testing.T) {
	list1 := StrList([]string{"a", "b", "c"})
	list2 := StrList([]string{"d", "e", "f"})
	list3 := StrList([]string{"a", "b", "c", "d", "e", "f"})
	if !reflect.DeepEqual(list3, list1.Union(list2)) {
		t.Error("stringlist Union error, list1 union list2 should be equal to list3")
	}
}

func TestStrListIntersection(t *testing.T) {
	l1 := StrList([]string{"a", "c", "e"})
	l2 := StrList([]string{"a", "b", "c", "d", "f"})
	l3 := StrList([]string{"b", "d", "f"})
	t.Log(l2.Intersection(l1))
	t.Log(l2.Intersection(l3))
	t.Log(l3.Intersection(l1))
	if !reflect.DeepEqual(l1.Intersection(l2), StrList([]string{"a", "c"})) {
		t.Error("stringlist Intersection error, intersect list should be [a,c]")
	}
	if !reflect.DeepEqual(l2.Intersection(l3), StrList([]string{"b", "d", "f"})) {
		t.Error("stringlist Intersection error, intersect list should be [b,d,f]")
	}
	if !reflect.DeepEqual(l3.Intersection(l1), StrList([]string{})) {
		t.Error("stringlist Intersection error, intersect list should be []")
	}
}

func TestStrListDifference(t *testing.T) {
	l1 := StrList([]string{"a", "c", "e"})
	l2 := StrList([]string{"a", "b", "c", "d", "f"})
	l3 := StrList([]string{"b", "d", "f"})
	t.Log(l2.Difference(l1))
	t.Log(l2.Difference(l3))
	t.Log(l3.Difference(l1))
	if !reflect.DeepEqual(l1.Difference(l2), StrList([]string{"e"})) {
		t.Error("stringlist Deference error, difference list should be [e]")
	}
	if !reflect.DeepEqual(l2.Difference(l3), StrList([]string{"a", "c"})) {
		t.Error("stringlist Deference error, difference list should be [a,c]")
	}
	if !reflect.DeepEqual(l3.Difference(l1), StrList([]string{"b", "d", "f"})) {
		t.Error("stringlist Deference error, difference list should be [b,d,f]")
	}
}

func TestStrListJoin(t *testing.T) {
	list := StrList([]string{"I", "am", "coding"})
	if list.Join(" ") != "I am coding" {
		t.Error("stringlist Join error, joined string is `I am coding`")
	}
}
