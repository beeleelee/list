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
// 	lee.Each(&intList, func(v lee.Item, i int){
// 		fmt.Println(v, i)
// 	})
// 	/*
// 	 *	0 0
// 	 *	1 1
// 	 *	2 2
// 	 */

// 	// list.Map
// 	intListMapped := lee.Map(&intList, func(v lee.Item, i int) lee.Item {
// 		return v.(int) * 2
// 	})

// 	fmt.Println(intListMapped.Data)
// 	// &{[0 2 4]}

// 	// list.Filter
// 	intListFiltered := lee.Filter(&intList, func(v lee.Item, i int) bool {
// 		return v.(int) % 2 == 1
// 	})

// 	fmt.Println(intListFiltered.Data)
// 	// &{[1]}
// }
// ```

//Item - generic type for list item
//
// in order to accept any type of item in collection
type Item interface{}

//Lister - interface for list
// Len return the size of the list
// Get return the item in the list by index
// Set return nil if successfully set item in the list by index,
//		return error if failed
// New return a new empty list
// Append item to extend the list with
type Lister interface {
	Len() int
	Get(int) (Item, error)
	Set(int, Item) error
	New(int) Lister
	Append(...Item)
}

// List a struct wrap collection in Data field
//
// why use struct, can I use slice, like []Item ?
// actually I use []Item at first time, but then
// went into trouble when implement Append method
type List struct {
	Data []Item
}

// Len return the length of the collection
func (l *List) Len() int {
	return len(l.Data)
}

// Get return the item in the collection by index
//
// if index in range, e will be nil
// out of range, e will has an error message
func (l *List) Get(i int) (item Item, e error) {
	length := l.Len()
	if i >= length {
		e = fmt.Errorf("List.Get index: %v is out of range", i)
		return
	}
	item = l.Data[i]
	e = nil
	return
}

// Set - modify the collection
//
// e will be nil if successfully set item
// e will has an message when index out of range
func (l *List) Set(i int, v Item) (e error) {
	size := l.Len()
	if i < 0 || i > size-1 {
		e = fmt.Errorf("the input index: %v is out of range", i)
		return
	}
	l.Data[i] = v
	e = nil
	return
}

// New - return an empty collection
func (*List) New(n int) Lister {
	r := new(List)
	r.Data = make([]Item, n)
	return r
}

// Append - extend the collection
func (l *List) Append(v ...Item) {
	l.Data = append(l.Data, v...)
}

func (l *List) Each(f EachFn) {
	Each(l, f)
}

func (l *List) Map(f MapFn) Lister {
	return Map(l, f)
}

func (l *List) Filter(f FilterFn) Lister {
	return Filter(l, f)
}

func (l *List) Equal(t Lister, f CmpFn) bool {
	return Equal(l, t, f)
}

func (l *List) FindIndex(f FilterFn) int {
	return FindIndex(l, f)
}

func (l *List) Find(f FilterFn) (Item, bool) {
	return Find(l, f)
}

func (l *List) Contains(f FilterFn) bool {
	return Contains(l, f)
}

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
		nl = List{data}
		e = nil
	} else {
		e = fmt.Errorf("ListFrom only accept slice or array input, but got %v", rv.Kind())
	}

	return
}

// Each - each loop
//
// use for loop to get item from list
// and feed item to EachFn
func Each(list Lister, f EachFn) {
	l := list.Len()
	for i := 0; i < l; i++ {
		item, _ := list.Get(i)
		f(item, i)
	}
}

// Map - map loop
//
// use for loop to get item from list
// and feed item to MapFn
func Map(list Lister, f MapFn) Lister {
	l := list.Len()
	mapedList := list.New(l)
	for i := 0; i < l; i++ {
		item, _ := list.Get(i)
		mapedList.Set(i, f(item, i))
	}
	return mapedList
}

// Filter - filter loop
//
// first create a new list by list.New
// then use each loop to get item from list
// and feed item to FilterFn which decide weather keep it or not
func Filter(list Lister, f FilterFn) Lister {
	filteredList := list.New(0)
	Each(list, func(v Item, i int) {
		if f(v, i) {
			filteredList.Append(v)
		}
	})
	return filteredList
}

// Equal - a way to compare whether two list is equal
//
// it accept a CmpFn which handle the equal logic
func Equal(s, t Lister, f CmpFn) (r bool) {
	sLen := s.Len()
	tLen := t.Len()
	r = true
	if sLen != tLen {
		r = false
		return
	}
	var sItem, tItem Item
	for i := 0; i < sLen; i++ {
		sItem, _ = s.Get(i)
		tItem, _ = t.Get(i)
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
func FindIndex(list Lister, f FilterFn) (index int) {
	l := list.Len()
	index = -1
	for i := 0; i < l; i++ {
		item, _ := list.Get(i)
		if f(item, i) {
			index = i
			break
		}
	}
	return
}

// Find - like FindIndex, but not return index of item
//
// it returns the specific item and ok flag
func Find(list Lister, f FilterFn) (r Item, ok bool) {
	l := list.Len()
	var item Item
	for i := 0; i < l; i++ {
		item, _ = list.Get(i)
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
func Contains(list Lister, f FilterFn) (r bool) {
	fmt.Println(Find(list, f))
	if _, ok := Find(list, f); ok {
		r = true
	} else {
		r = false
	}
	return
}
