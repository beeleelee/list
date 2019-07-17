package list

//Item - generic type for list item 
type Item interface {}

//Lister - interface for list 
type Lister interface {
	Len() int 
	Get(i int) Item 
	Set(i int, v Item)
	New(n int) Lister
	Append(v Item)
}

//EachFn  type for each function handle 
type EachFn func(v Item, i int)

//MapFn type for map function handle
type MapFn func(v Item, i int) Item 

//FilterFn type for filter function handle
type FilterFn func(v Item, i int) bool 

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
	mapedList := list.New(l) 
	for i := 0; i < l; i++ {
		mapedList.Set(i, f(list.Get(i), i))
	}
	return mapedList
}

func Filter(list Lister, f FilterFn) Lister {
	filteredList := list.New(0)
	Each(list, func(v Item, i int) {
		if f(v, i) {
			filteredList.Append(v)
		}
	})
	return filteredList
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

type List struct {
	Data []Item 
}

func (l *List) Len() int {
	return len(l.Data)
}

func (l *List) Get(i int) Item {
	return l.Data[i]
}

func (l *List) Set(i int, v Item) {
	l.Data[i] = v 
}

func (_ *List) New(n int) Lister {
	r := new(List)
	r.Data = make([]Item, n)
	return r
}

func (l *List) Append(v Item) {
	l.Data = append(l.Data, v)
}