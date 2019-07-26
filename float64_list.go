package list

import (
	"math"
	"math/rand"
	"sort"
	"time"
)

// F64List implements Each Map Filter ... for float64 slice
// for the sake of better performance
type F64List []float64

// F64LEachFn method Each handler signature
type F64LEachFn func(v float64, i int)

// F64LMapFn method Map handler signature
type F64LMapFn func(v float64, i int) float64 

// F64LItemTestFn method Filter handler signature
type F64LItemTestFn func(v float64, i int) bool

// F64LReduceFn method Reduce handler signature
type F64LReduceFn func(a, b float64) float64 

func (list F64List) Each(f F64LEachFn) F64List {
	for i, v := range list {
		f(v, i)
	}
	return list
}

func (list F64List) Map(f F64LMapFn) (r F64List) {
	r = F64List(make([]float64, len(list)))
	for i, v := range list {
		r[i] = f(v, i)
	}
	return
}

func (list F64List) Filter(f F64LItemTestFn) (r F64List) {
	r = F64List([]float64{})
	for i, v := range list {
		if f(v, i) {
			r = append(r, v)
		}
	}	
	return 
}

func (list F64List) FindIndex(v float64) (index int) {
	l := len(list)
	index = -1
	for i := 0; i < l; i++ {
		if list[i] == v {
			index = i 
			break
		}
	}
	return 
}

func (list F64List) Contains(v float64) (r bool) {
	if list.FindIndex(v) > -1 {
		r = true
	}else{
		r = false
	}
	return
}

func (list F64List) Reduce(f F64LReduceFn, startv float64) (r float64) {
	r = startv 
	for _, v := range list {
		r = f(r, v)
	}
	return 
}

func (list F64List) Some(f F64LItemTestFn) (r bool) {
	l := len(list)
	r = false 
	for i := 0; i < l; i++ {
		if f(list[i], i) {
			r = true
			break
		}
	}
	return 
}

func (list F64List) Every(f F64LItemTestFn) (r bool) {
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

func (list F64List) Len() int {
	return len(list)
}

func (list F64List) Less(i, j int) bool {
	return list[i] < list[j]
}

func (list F64List) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func (list F64List) Shuffle() (r F64List) {
	l := len(list)
	r = make([]float64, l)
	copy(r, list)
	if l > 1 {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(l, func(i, j int) {
			r[i], r[j] = r[j], r[i]
		})
	}
	return
}

func (list F64List) Sort() {
	sort.Sort(list)
}

func (list F64List) IsSorted() bool {
	return sort.IsSorted(list)
}

func (list F64List) Get(i int) (float64, bool) {
	l := len(list)
	if l == 0 {
		return 0, false
	}
	index := i
	if i > l-1 {
		return 0, false
	}
	if i < 0 {
		index = l + i
	}
	if index < 0 {
		return 0, false
	}
	return list[index], true
}

func (list F64List) Tail(n int) F64List {
	l := len(list)
	if l == 0 || n >= l {
		return list
	}

	return list[l-n:]
}

func (list F64List) Union(t F64List) F64List {
	return append(list, t...)
}

func (list F64List) Intersection(t F64List) (r F64List) {
	minLen := int(math.Min(float64(len(list)), float64(len(t))))
	r = F64List(make([]float64, minLen))
	if minLen == 0 {
		return
	}
	index := 0
	for _, v := range list {
		if t.Contains(v) {
			r[index] = v
			index++
		}
	}
	r = r[0:index]
	return
}

func (list F64List) Difference(t F64List) (r F64List) {
	minLen := int(math.Min(float64(len(list)), float64(len(t))))
	r = F64List(make([]float64, minLen))
	if minLen == 0 {
		return
	}
	index := 0
	for _, v := range list {
		if !t.Contains(v) {
			r[index] = v
			index++
		}
	}
	r = r[0:index]
	return
}