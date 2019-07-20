package list

import (
	"fmt"
	"reflect"
)

// #package list
//
//	list provide some useful utilities
//
// 	in order to manipulate collection conveniently
//
//	in the form of functional programming
//
//	hope it will be helpful
//
// ##Example

// ```
// package main

// import (
// 	"fmt"
// 	lee "github.com/beeleelee/list"
// )

// func main() {
// 	intList, _ := lee.From([]int{0,1,2})
// 	// list.Each
// 	intList.Each(func(v lee.Item, i int){
// 		fmt.Println(v, i)
// 	})
// 	// lee.Each(intList, func(v lee.Item, i int){
// 	// 	fmt.Println(v, i)
// 	// })

// 	/*
// 	 *	0 0
// 	 *	1 1
// 	 *	2 2
// 	 */

// 	// list.Map
// 	intListMapped := intList.Map(func(v lee.Item, i int) lee.Item {
// 		return v.(int) * 2
// 	})
// 	// intListMapped := lee.Map(intList, func(v lee.Item, i int) lee.Item {
// 	// 	return v.(int) * 2
// 	// })

// 	fmt.Println(intListMapped)
// 	// [0 2 4]

// 	// list.Filter
// 	intListFiltered := intList.Filter(func(v lee.Item, i int) bool {
// 		return v.(int) % 2 == 1
// 	})

// 	// intListFiltered := lee.Filter(intList, func(v lee.Item, i int) bool {
// 	// 	return v.(int) % 2 == 1
// 	// })

// 	fmt.Println(intListFiltered)
// 	// [1]
// }
// ```

//Item - generic type for list item
//
// in order to accept any type of item in collection
type Item interface{}

// List a struct wrap collection in Data field
//
type List []Item

//EachFn  each loop handle signature
//
// func(v Item, i int){
// 	// switch value to the expected type
// 	sv, _ := v.(int) // just for example, actually can use any type you specified
// 	fmt.Println(sv)
// }
type EachFn func(Item, int)

//MapFn  map loop handle signature
//
// func(v Item, i int) (item Item) {
// 	sv, _ := v.(float64)
// 	return sv * sv
// }
type MapFn func(Item, int) Item

//FilterFn filter loop handle signature
//
// func(v Item, i int) bool {
// 	sv := v.(string)
// 	return sv == "foo"
// }
type FilterFn func(Item, int) bool

//CmpFn compare handle signature
//
// func(a, b Item) bool {
// 	return a == b
// }
type CmpFn func(a, b Item) bool

// ReduceFn reduce handle signature
type ReduceFn func(a, b Item) Item

// Each convenience wrapper for Each function
func (l List) Each(f EachFn) List {
	Each(l, f)
	return l 
}
// Map convenience wrapper for Map function
func (l List) Map(f MapFn) List {
	return Map(l, f)
}
// Filter convenience wrapper for Filter function
func (l List) Filter(f FilterFn) List {
	return Filter(l, f)
}
// Equal convenience wrapper for Equal function
func (l List) Equal(t List, f CmpFn) bool {
	return Equal(l, t, f)
}
// FindIndex convenience wrapper for FindIndex function
func (l List) FindIndex(f FilterFn) int {
	return FindIndex(l, f)
}
// Find convenience wrapper for Find function
func (l List) Find(f FilterFn) (Item, bool) {
	return Find(l, f)
}
// Contains convenience wrapper for Contains function
func (l List) Contains(f FilterFn) bool {
	return Contains(l, f)
}
// Reduce convenience wrapper for Reduce function
func (l List) Reduce(f ReduceFn, a Item) Item {
	return Reduce(l, f, a)
}



//From - convert regular slice to List
//
//	as do not know the item type in the slic
// 	so use reflect package to get the item type
//	and rebuild a new slice with Item type
//
//	call it like this:
// 	list.From([]int{1,2,3})
func From(source interface{}) (nl List, e error) {
	rv := reflect.ValueOf(source)
	if rv.Kind() == reflect.Slice || rv.Kind() == reflect.Array {
		rvLen := rv.Len()
		data := make([]Item, rvLen)
		for i := 0; i < rvLen; i++ {
			data[i] = rv.Index(i).Interface()
		}
		nl = data
		e = nil
	} else {
		e = fmt.Errorf("ListFrom only accept slice or array input, but got %v", rv.Kind())
	}
	return
}

func FromInts(source []int) (nl List) {
	nl = New(len(source))
	for i, v := range source {
		nl[i] = v
	}
	return
}

// New generate a new List instance
func New(length int) List {
	return make([]Item, length)
}

// Each - each loop
//
// use for loop to get item from list
// and feed item to EachFn
func Each(list List, f EachFn) {
	for i, v := range list {
		f(v, i)
	}
}

// Map - map loop
//
// use for loop to get item from list
// and feed item to MapFn
func Map(list List, f MapFn) List {
	l := len(list)
	mapedList := make([]Item, l)
	for i, v := range list {
		mapedList[i] = f(v, i)
	}
	return mapedList
}

// Filter - filter loop
//
// first create a new list 
// then use each loop to get item from list
// and feed item to FilterFn which decide weather keep it or not
func Filter(list List, f FilterFn) List {
	filteredList := make([]Item, 0)
	Each(list, func(v Item, i int) {
		if f(v, i) {
			filteredList = append(filteredList, v)
		}
	})
	return filteredList
}

// Equal - a way to compare whether two list is equal
//
// it accept a CmpFn which handle the equal logic
func Equal(s, t List, f CmpFn) (r bool) {
	sLen := len(s)
	tLen := len(t)
	r = true
	if sLen != tLen {
		r = false
		return
	}
	var sItem, tItem Item
	for i := 0; i < sLen; i++ {
		sItem = s[i]
		tItem = t[i]
		if !f(sItem, tItem) {
			r = false
			break
		}
	}
	return
}

// FindIndex - a way to find the index of a specific item
//
//	it return -1 if could not find the item
//	it accept a FilterFn which will specific the item
func FindIndex(list List, f FilterFn) (index int) {
	l := len(list)
	index = -1
	for i := 0; i < l; i++ {
		if f(list[i], i) {
			index = i
			break
		}
	}
	return
}

// Find - like FindIndex, but not return index of item
//
// it returns the specific item and ok flag
func Find(list List, f FilterFn) (r Item, ok bool) {
	l := len(list)
	var item Item
	for i := 0; i < l; i++ {
		item = list[i]
		ok = false
		if f(item, i) {
			r = item
			ok = true
			break
		}
	}
	return
}

// Contains - like Find
//
// return true if find the item
// return false if can not find the item
func Contains(list List, f FilterFn) (r bool) {
	if _, ok := Find(list, f); ok {
		r = true
	} else {
		r = false
	}
	return
}

// Reduce - fold the list
func Reduce(list List, f ReduceFn, a Item) (r Item) {
	l, i := len(list), 0
	if a == nil { // use first item to start if not pass a start value
		a = list[i]
		i++
	}
	for ; i < l; i++ {
		a = f(a, list[i])
	}
	r = a
	return
}
