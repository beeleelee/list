package list_test

import (
	"testing"
	"reflect"
	. "../list"
)

func cmp(a, b Item) bool {
	return a == b
}

func TestEach(t *testing.T) {
	list, err := From([]int{1,3,5,7,9})
	if err != nil {
		t.Errorf("%v", err)
		return 
	}
	list2 := New(len(list))
	Each(list, func(v Item, i int){
		list2[i] = v
	})
	t.Log(list, list2)
	if !reflect.DeepEqual(list, list2) {
		t.Errorf("expect list: %v and list2: %v to be equal", list, list2)
	}
}

func TestMap(t *testing.T) {
	list := FromInts([]int{1,3,5,7,9})
	list2 := list.Map(func(v Item, i int) Item {
		return v.(int) * 2
	})
	if !Equal(list2, []Item{2,6,10,14,18}, cmp) {
		t.Errorf("expect list2: %v to be %v", list2, []Item{2,6,10,14,18})
	}
}

func TestFilter(t *testing.T) {
	list, _ := From([]int{1,2,3,4,5,6,7})
	list2 := list.Filter(func(v Item, i int) bool {
		return v.(int) % 2 == 0
	})
	if !Equal(list2, []Item{2,4,6}, cmp) {
		t.Errorf("expect list2: %v to be %v", list2, []Item{2,4,6})
	}
}

func TestFindIndex(t *testing.T) {
	list, _ := From([]rune{'a', 'b', 'c', 'd', 'e'})
	index := list.FindIndex(func(v Item, i int) bool {
		return v.(rune) == 'e'
	})
	if index != 4 {
		t.Errorf("expect item index is 4 but got %v", index)
	}
}

func TestEqual(t *testing.T) {
	a := []Item{'a', 'b', 'c', 'd', 'e'}
	b := []Item{'a', 'b', 'c', 'd', 'e'}
	
	if !Equal(a, b, cmp) {
		t.Errorf("expect a: %v should be equal to b: %v, but got not equal", a, b)
	}
}

func TestFind(t *testing.T) {
	type User struct {
		name string
		age int
	}
	userAlex := User{"alex", 38}
	userBeeleelee := User{"beeleelee", 40}
	list, _ := From([]User{userAlex,userBeeleelee})
	item, ok := Find(list, func(v Item, i int) bool {
		return v.(User).name == "alex"
	})
	
	if userAlex == item {
		t.Logf("%v equals %v, %v", item, userAlex, ok)
	}
	
	if !ok {
		t.Errorf("find error, it should find user alex")
	}

	_, ok = Find(list, func(v Item, i int) bool {
		return v.(User).name == "foo"
	})

	if ok {
		t.Errorf("find error, how could it find user foo")
	}
}

func TestContains(t *testing.T) {
	type User struct {
		name string
		age int
	}
	list, _ := From([]User{{"alex", 38},{"beeleelee", 40}})
	hasFoo := Contains(list, func(v Item, i int) bool {
		return v.(User).name == "foo"
	})
	
	if hasFoo {
		t.Errorf("contains error, list does not contains user foo")
	}
}

func TestReduce(t *testing.T) {
	intList, _ := From([]int{0,1,2,3,4,5})
	total := Reduce(intList, func(a, b Item) Item {
		return a.(int) + b.(int)
	}, nil)
	if total.(int) != 15 {
		t.Errorf("reduce error, expect total to be 15 but got %v", total)
	}
	sum20 := intList.Reduce(func(a, b Item) Item {
		return a.(int) + b.(int)
	}, 20)
	if sum20 != 35 {
		t.Errorf("reduce error, expect sum20 to be 35 but got %v", sum20)
	}
}

func TestSome(t *testing.T) {
	list := FromStrings([]string{"apple", "orange", "banana"})
	hasBanana := Some(list, func(v Item, i int) bool {
		return v.(string) == "banana"
	})
	t.Log("has banana: ", hasBanana)
	if !hasBanana {
		t.Errorf("some error, it should has banana")
	}
	hasFoo := list.Some(func(v Item, i int) bool {
		return v.(string) == "foo"
	})
	t.Log("has banana: ", hasFoo)
	if hasFoo {
		t.Error("some error, it should not has foo")
	}
}

func TestEvery(t *testing.T) {
	list := FromFloat64s([]float64{2.0, 3.0, 4.0})
	biggerThan1 := Every(list, func(v Item, i int) bool {
		return v.(float64) > 1
	})
	t.Log("bigger than 1: ", biggerThan1)
	if !biggerThan1 {
		t.Errorf("every error, all items should bigger than 1")
	}
	biggerThan3 := list.Every(func(v Item, i int) bool {
		return v.(float64) > 3
	})
	t.Log("bigger than 3: ", biggerThan3)
	if biggerThan3 {
		t.Error("every error, some item should not bigger than 3")
	}
}

func TestShuffle(t *testing.T) {
	list := FromInts([]int{0,1,2,3,4,5,6,7,8,9})
	list2 := Shuffle(list)
	t.Log(list2)
	list3 := list.Shuffle()
	t.Log(list3)
}

func TestSort(t *testing.T) {
	list := FromInts([]int{9,5,3,7,2,8,0,4,1,6})
	Sort(list, func(i, j int) bool {
		return list[i].(int) < list[j].(int)
	})
	t.Log(list)
	if !reflect.DeepEqual(list, FromInts([]int{0,1,2,3,4,5,6,7,8,9})) {
		t.Error("sort error, sorted items has wrong order")
	}
	type Student struct {
		name string 
		id int 
	}
	list2, _ := From([]Student{{"leo", 1}, {"lucy", 2}, {"tom", 3}})
	Sort(list2, func(i, j int) bool {
		ui := list2[i].(Student)
		uj := list2[j].(Student)
		return ui.id > uj.id
	})
	t.Log(list2)
	if list2[0].(Student).name != "tom" {
		t.Error("sort error, the first item of sorted list2 should be tom")
	}
}