package list

type IntList []int

type ILEachFn func(v, i int)

type ILMapFn func(v, i int) int 

type ILItemTestFn func(v, i int) bool

func (l IntList) Each(f ILEachFn) IntList {
	for i, v := range l {
		f(v, i)
	}
	return l
}

func (l IntList) Map(f ILMapFn) (r IntList) {
	r = IntList(make([]int, len(l)))
	for i, v := range l {
		r[i] = f(v, i)
	}
	return
}

func (l IntList) Filter(f ILItemTestFn) (r IntList) {
	r = IntList([]int{})
	for i, v := range l {
		if f(v, i) {
			r = append(r, v)
		}
	}	
	return 
}