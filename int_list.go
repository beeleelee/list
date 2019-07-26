package list

import (
	//"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

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

func (list IntList) Each(f ILEachFn) IntList {
	for i, v := range list {
		f(v, i)
	}
	return list
}

func (list IntList) Map(f ILMapFn) (r IntList) {
	r = IntList(make([]int, len(list)))
	for i, v := range list {
		r[i] = f(v, i)
	}
	return
}

func (list IntList) Filter(f ILItemTestFn) (r IntList) {
	r = IntList([]int{})
	for i, v := range list {
		if f(v, i) {
			r = append(r, v)
		}
	}	
	return 
}

func (list IntList) FindIndex(v int) (index int) {
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

func (list IntList) Contains(v int) (r bool) {
	if list.FindIndex(v) > -1 {
		r = true
	}else{
		r = false
	}
	return
}

func (list IntList) Reduce(f ILReduceFn, startv int) (r int) {
	r = startv 
	for _, v := range list {
		r = f(r, v)
	}
	return 
}

func (list IntList) Some(f ILItemTestFn) (r bool) {
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

func (list IntList) Every(f ILItemTestFn) (r bool) {
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

func (list IntList) Len() int {
	return len(list)
}

func (list IntList) Less(i, j int) bool {
	return list[i] < list[j]
}

func (list IntList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func (list IntList) Shuffle() (r IntList) {
	l := len(list)
	r = make([]int, l)
	copy(r, list)
	if l > 1 {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(l, func(i, j int) {
			r[i], r[j] = r[j], r[i]
		})
	}
	return
}

func (list IntList) Sort() {
	sort.Sort(list)
}

func (list IntList) IsSorted() bool {
	return sort.IsSorted(list)
}

func (list IntList) Get(i int) (int, bool) {
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

func (list IntList) Tail(n int) IntList {
	l := len(list)
	if l == 0 || n >= l {
		return list
	}

	return list[l-n:]
}

func (list IntList) Union(t IntList) IntList {
	return append(list, t...)
}

func (list IntList) Intersection(t IntList) (r IntList) {
	minLen := int(math.Min(float64(len(list)), float64(len(t))))
	r = IntList(make([]int, minLen))
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