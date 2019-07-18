package list

import (
	"fmt"
	"reflect"
)

// #package list

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
	Get(int) Item
	Set(int, Item) error
	New(int) Lister
	Append(...Item)
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
		f(list.Get(i), i)
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
		mapedList.Set(i, f(list.Get(i), i))
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

func Equal(s, t Lister, f CmpFn) (r bool) {
	sLen := s.Len()
	tLen := t.Len()
	r = true
	if sLen != tLen {
		r = false
		return
	}
	for i := 0; i < sLen; i++ {
		if !f(s.Get(i), t.Get(i)) {
			r = false
			break
		}
	}
	return
}

func FindIndex(list Lister, f FilterFn) (index int) {
	l := list.Len()
	index = -1
	for i := 0; i < l; i++ {
		if f(list.Get(i), i) {
			index = i
			break
		}
	}
	return
}

func Find(list Lister, f FilterFn) (r Item) {
	l := list.Len()
	r = nil
	for i := 0; i < l; i++ {
		r = list.Get(i)
		if f(r, i) {
			break
		}
	}
	return
}

type List struct {
	Data []Item
}

func (l *List) Len() int {
	return len(l.Data)
}

func (l *List) Get(i int) Item {
	return l.Data[i]
}

func (l *List) Set(i int, v Item) (e error) {
	size := l.Len()
	if i < 0 || i > size-1 {
		e = fmt.Errorf("*List Set - the input index: %v is out of range\n", i)
		return
	}
	l.Data[i] = v
	e = nil
	return
}

func (_ *List) New(n int) Lister {
	r := new(List)
	r.Data = make([]Item, n)
	return r
}

func (l *List) Append(v ...Item) {
	l.Data = append(l.Data, v...)
}
