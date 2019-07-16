package list

// generic type for list item 
type Item interface {}

// interface for list 
type Lister interface {
	Len() int 
	Get(i int) Item 
	Set(i int, v Item)
	New() Lister 
}

// type for each function handle 
type EachFn func(v Item, i int, l Lister)

// type for map function handle
type MapFn func(v Item, i int, l Lister) Item 

