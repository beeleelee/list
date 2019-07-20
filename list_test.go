package list_test

import (
	"fmt"
	"testing"
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
	list2 := New(0)
	Each(list, func(v Item, i int){
		list2 = append(list2, v)
	})
	if !Equal(list, list2, cmp) {
		t.Errorf("seems Each not works")
	}
	fmt.Println("...")
}

func TestMap(t *testing.T) {
	list, _ := From([]int{1,3,5,7,9})
	newList := list.Map(func(v Item, i int) Item {
		return v.(int) * 2
	})
	if !Equal(newList, []Item{2,6,10,14,18}, cmp) {
		t.Errorf("before map %v \n after map %v", list, newList)
	}
}

// func TestFilter(t *testing.T) {
// 	list, _ := From([]int{1,2,3,4,5,6,7})
// 	list2 := Filter(list, func(v Item, i int) bool {
// 		return v.(int) % 2 == 0
// 	})
// 	if !Equal(list2, &List{[]Item{2,4,6}}, cmp) {
// 		t.Errorf("Filter seems not work")
// 	}
// }

// func TestFindIndex(t *testing.T) {
// 	list, _ := From([]rune{'a', 'b', 'c', 'd', 'e'})
// 	index := FindIndex(list, func(v Item, i int) bool {
// 		return v.(rune) == 'e'
// 	})
// 	if index != 4 {
// 		t.Errorf("expect item index is 4 but got %v", index)
// 	}
// }

// func TestEqual(t *testing.T) {
// 	a := List{[]Item{'a', 'b', 'c', 'd', 'e'}}
// 	b := List{[]Item{'a', 'b', 'c', 'd', 'e'}}
	
// 	if !Equal(&a, &b, cmp) {
// 		t.Errorf("s %v should be equal to t %v, but got not equal", a, b)
// 	}
// }

// func TestFind(t *testing.T) {
// 	type User struct {
// 		name string
// 		age int
// 	}
// 	userAlex := User{"alex", 38}
// 	userBeeleelee := User{"beeleelee", 40}
// 	list, _ := From([]User{userAlex,userBeeleelee})
// 	item, ok := Find(list, func(v Item, i int) bool {
// 		return v.(User).name == "alex"
// 	})
	
// 	if userAlex == item {
// 		fmt.Printf("%v equals %v, %v", item, userAlex, ok)
// 	}else{
// 		fmt.Printf("%v not equals %v, %v", userAlex, item, ok)
// 	}
	
// 	if !ok {
// 		t.Errorf("seems Find not work")
// 	}
// }

// func TestContains(t *testing.T) {
// 	type User struct {
// 		name string
// 		age int
// 	}
// 	list, _ := From([]User{{"alex", 38},{"beeleelee", 40}})
// 	hasFoo := Contains(list, func(v Item, i int) bool {
// 		return v.(User).name == "foo"
// 	})
	
// 	if hasFoo {
// 		t.Errorf("seems Contain not work")
// 	}
// }

// func TestReduce(t *testing.T) {
// 	intList, _ := From([]int{0,1,2,3,4,5})
// 	total := Reduce(intList, func(a, b Item) Item {
// 		return a.(int) + b.(int)
// 	}, nil)
// 	if total.(int) != 15 {
// 		t.Errorf("Reduce not work, expect total to be 15 but got %v", total)
// 	}
// 	sum20 := intList.Reduce(func(a, b Item) Item {
// 		return a.(int) + b.(int)
// 	}, 20)
// 	if sum20 != 35 {
// 		t.Errorf("Reduce not work, expect sum20 to be 35 but got %v", sum20)
// 	}
// }