package list

//Item - generic type for list item 
type Item interface {}

//Lister - interface for list 
type Lister interface {
	Len() int 
	Get(i int) Item 
	Set(i int, v Item)
}

//EachFn  type for each function handle 
type EachFn func(v Item, i int)

//MapFn type for map function handle
type MapFn func(v Item, i int) Item 

// Each - each loop handler
func Each(list Lister, f EachFn){
	l := list.Len()
	for i := 0; i < l; i++ {
		f(list.Get(i), i)
	}
}

// Map - map loop handler
func Map(list Lister, f MapFn) Lister {
	l := list.Len()
	mapedList := list 
	for i := 0; i < l; i++ {
		mapedList.Set(i, f(list.Get(i), i))
	}
	return mapedList
}

type List struct {
	Data []Item 
}

func (l List) Len() int {
	return len(l.Data)
}

func (l List) Get(i int) Item {
	return l.Data[i]
}

func (l List) Set(i int, v Item) {
	l.Data[i] = v 
}