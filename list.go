package list

// generic type for list item 
type Item interface {}

// interface for list 
type Lister interface {
	Len() int 
	Get(i int) Item 
	Set(i int, v Item)
}

// type for each function handle 
type EachFn func(v Item, i int)

// type for map function handle
type MapFn func(v Item, i int) Item 

func Each(list Lister, f EachFn){
	l := list.Len()
	for i := 0; i < l; i++ {
		f(list.Get(i), i)
	}
}

func Map(list Lister, f MapFn) Lister {
	l := list.Len()
	mapedList := list 
	for i := 0; i < l; i++ {
		mapedList.Set(i, f(list.Get(i), i))
	}
	return mapedList
}