package list

// IntList implements Each Map Filter ... for int slice
// for the sake of better performance
type IntList []int

// ILEachFn method Each handler signature
type ILEachFn func(v, i int)

// ILMapFn method Map handler signature
type ILMapFn func(v, i int) int 

// ILItemTestFn method Filter handler signature
type ILItemTestFn func(v, i int) bool

// ILReduceFn method Reduce handler signature
type ILReduceFn func(a, b int) int 

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

func (l IntList) Contains(v int) (r bool) {
	if l.FindIndex(v) > -1 {
		r = true
	}else{
		r = false
	}
	return
}

func (l IntList) Reduce(f ILReduceFn, startv int) (r int) {
	r = startv 
	for _, v := range l {
		r = f(r, v)
	}
	return 
}

func (l IntList) Some(f ILItemTestFn) (r bool) {
	size := len(l)
	r = false 
	for i := 0; i < size; i++ {
		if f(l[i], i) {
			r = true
			break
		}
	}
	return 
}

func (l IntList) Every(f ILItemTestFn) (r bool) {
	size := len(l)
	r = true 
	for i := 0; i < size; i++ {
		if !f(l[i], i) {
			r = false
			break
		}
	}
	return 
}

func (l IntList) Len() int {
	return len(l)
}

func (l IntList) Less(i, j int) bool {
	return l[i] < l[j]
}

func (l IntList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}
