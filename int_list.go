package list

// IntList implements Each Map Filter ... for int slice
// for the sake of better performance
type IntList []int

// ILEachFn method Each handle signature
type ILEachFn func(v, i int)

// ILMapFn method Map handle signature
type ILMapFn func(v, i int) int 

// ILItemTestFn method Filter handle signature
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

func (l IntList) FindIndex(v int) (index int) {
	length := len(l)
	index = -1
	for i := 0; i < length; i++ {
		if l[i] == v {
			index = i 
			break
		}
	}
	return 
}