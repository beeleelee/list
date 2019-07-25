package list

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
	"sort"
	"time"
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

//ItemTestFn filter loop handle signature
//
// func(v Item, i int) bool {
// 	sv := v.(string)
// 	return sv == "foo"
// }
type ItemTestFn func(Item, int) bool

//EqualFn compare handle signature
//
// func(a, b Item) bool {
// 	return a == b
// }
type EqualFn func(a, b Item) bool

// ReduceFn reduce handle signature
type ReduceFn func(a, b Item) Item

// SwapFn swap items by index
type SwapFn func(i, j int)

// LessFn same signature as sort.Less
type LessFn func(i, j int) bool

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
func (l List) Filter(f ItemTestFn) List {
	return Filter(l, f)
}

// Equal convenience wrapper for Equal function
func (l List) Equal(t List, f EqualFn) bool {
	return Equal(l, t, f)
}

// FindIndex convenience wrapper for FindIndex function
func (l List) FindIndex(f ItemTestFn) int {
	return FindIndex(l, f)
}

// Find convenience wrapper for Find function
func (l List) Find(f ItemTestFn) (Item, bool) {
	return Find(l, f)
}

// Contains convenience wrapper for Contains function
func (l List) Contains(f ItemTestFn) bool {
	return Contains(l, f)
}

// Reduce convenience wrapper for Reduce function
func (l List) Reduce(f ReduceFn, a Item) Item {
	return Reduce(l, f, a)
}

// Some convenience wrapper for Some Function
func (l List) Some(f ItemTestFn) bool {
	return Some(l, f)
}

// Every convenience wrapper for Every Function
func (l List) Every(f ItemTestFn) bool {
	return Every(l, f)
}

// Shuffle convenience wrapper for Shuffle Function
func (l List) Shuffle() List {
	return Shuffle(l)
}

// Sort convenience wrapper for std sort.Slice
func (l List) Sort(f LessFn) List {
	sort.Slice(l, f)
	return l
}

// IsSorted convenience wrapper for std sort.SliceIsSorted
func (l List) IsSorted(f LessFn) bool {
	return sort.SliceIsSorted(l, f)
}

// Get convenience wrapper for Get Function
func (l List) Get(i int) Item {
	return Get(l, i)
}

// Tail convenience wrapper for Tail Function
func (l List) Tail(n int) List {
	return Tail(l, n)
}

// Union convenience wrapper for Union Function
func (l List) Union(t List) List {
	return Union(l, t)
}

// Intersection convenience wrapper for Intersection Function
func (l List) Intersection(t List, f EqualFn) List {
	return Intersection(l, t, f)
}

// Difference convenience wrapper for Difference Function
func (l List) Difference(t List, f EqualFn) List {
	return Difference(l, t, f)
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

// FromInts convert int slice to List
func FromInts(source []int) (nl List) {
	nl = New(len(source))
	for i, v := range source {
		nl[i] = v
	}
	return
}

// FromFloat64s convert float64 slice to List
func FromFloat64s(source []float64) (nl List) {
	nl = New(len(source))
	for i, v := range source {
		nl[i] = v
	}
	return
}

// FromStrings convert string slice to List
func FromStrings(source []string) (nl List) {
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
// and feed item to ItemTestFn which decide weather keep it or not
func Filter(list List, f ItemTestFn) List {
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
// it accept a EqualFn which handle the equal logic
func Equal(s, t List, f EqualFn) (r bool) {
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
//	it accept a ItemTestFn which will specific the item
func FindIndex(list List, f ItemTestFn) (index int) {
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
func Find(list List, f ItemTestFn) (r Item, ok bool) {
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
func Contains(list List, f ItemTestFn) (r bool) {
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

// Some - return true if any item pass test
func Some(list List, f ItemTestFn) (r bool) {
	l := len(list)
	for i := 0; i < l; i++ {
		if f(list[i], i) {
			r = true
			break
		}
	}
	return
}

// Every - return true if every item pass test
func Every(list List, f ItemTestFn) (r bool) {
	l := len(list)
	r = true
	for i := 0; i < l; i++ {
		if !f(list[i], i) {
			r = false
			break
		}
	}
	return
}

// Shuffle - return a shuffled list
func Shuffle(list List) (r List) {
	l := len(list)
	r = make([]Item, l)
	copy(r, list)
	if l > 1 {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(l, func(i, j int) {
			r[i], r[j] = r[j], r[i]
		})
	}
	return
}

// Sort - convenience wrapper for std sort.Slice
func Sort(list List, f LessFn) {
	sort.Slice(list, f)
}

// IsSorted - convenience wrapper for std sort.SliceIsSorted
func IsSorted(list List, f LessFn) bool {
	return sort.SliceIsSorted(list, f)
}

// Get - get item from list
// it can accept negative int as index, like -1
// attention: it will never failed
// if then index out of range, or no item in list,  it will return nil
func Get(list List, i int) Item {
	l := len(list)
	if l == 0 {
		return nil
	}
	index := i
	if i > l-1 {
		return nil
	}
	if i < 0 {
		index = l + i
	}
	if index < 0 {
		return nil
	}
	return list[index]
}

// Tail - get items from last
func Tail(list List, n int) List {
	l := len(list)
	if l == 0 || n >= l {
		return list
	}

	return list[l-n:]
}

// Union - union two lists
func Union(s List, t List) List {
	return append(s, t...)
}

// Intersection - return a list with items in both list
func Intersection(s List, t List, f EqualFn) (r List) {
	minLen := int(math.Min(float64(len(s)), float64(len(t))))
	r = List(make([]Item, minLen))
	if minLen == 0 {
		return
	}
	index := 0
	for _, v := range s {
		if Contains(t, func(sv Item, _ int) bool {
			return f(v, sv)
		}) {
			r[index] = v
			index++
		}
	}
	r = r[0:index]
	return
}

// Difference - return a list with items not in the other list
func Difference(s List, t List, f EqualFn) (r List) {
	minLen := int(math.Min(float64(len(s)), float64(len(t))))
	r = List(make([]Item, minLen))
	if minLen == 0 {
		return
	}
	index := 0
	for _, v := range s {
		if !Contains(t, func(sv Item, _ int) bool {
			return f(v, sv)
		}) {
			r[index] = v
			index++
		}
	}
	r = r[0:index]
	return
}
