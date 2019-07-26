package list

import (
	"math"
	"math/rand"
	"sort"
	"time"
)

// StrList implements Each Map Filter ... for string slice
// for the sake of better performance
type StrList []string

// StrLEachFn method Each handler signature
type StrLEachFn func(v string, i int)

// StrLMapFn method Map handler signature
type StrLMapFn func(v string, i int) string

// StrLItemTestFn method Filter handler signature
type StrLItemTestFn func(v string, i int) bool

func (list StrList) Each(f StrLEachFn) StrList {
	for i, v := range list {
		f(v, i)
	}
	return list
}

func (list StrList) Map(f StrLMapFn) (r StrList) {
	r = StrList(make([]string, len(list)))
	for i, v := range list {
		r[i] = f(v, i)
	}
	return
}

func (list StrList) Filter(f StrLItemTestFn) (r StrList) {
	r = StrList([]string{})
	for i, v := range list {
		if f(v, i) {
			r = append(r, v)
		}
	}
	return
}

func (list StrList) FindIndex(v string) (index int) {
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

func (list StrList) Contains(v string) (r bool) {
	if list.FindIndex(v) > -1 {
		r = true
	} else {
		r = false
	}
	return
}

func (list StrList) Some(f StrLItemTestFn) (r bool) {
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

func (list StrList) Every(f StrLItemTestFn) (r bool) {
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

func (list StrList) Len() int {
	return len(list)
}

func (list StrList) Less(i, j int) bool {
	return list[i] < list[j]
}

func (list StrList) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func (list StrList) Shuffle() (r StrList) {
	l := len(list)
	r = make([]string, l)
	copy(r, list)
	if l > 1 {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(l, func(i, j int) {
			r[i], r[j] = r[j], r[i]
		})
	}
	return
}

func (list StrList) Sort() {
	sort.Sort(list)
}

func (list StrList) IsSorted() bool {
	return sort.IsSorted(list)
}

func (list StrList) Get(i int) (string, bool) {
	l := len(list)
	if l == 0 {
		return "", false
	}
	index := i
	if i > l-1 {
		return "", false
	}
	if i < 0 {
		index = l + i
	}
	if index < 0 {
		return "", false
	}
	return list[index], true
}

func (list StrList) Tail(n int) StrList {
	l := len(list)
	if l == 0 || n >= l {
		return list
	}

	return list[l-n:]
}

func (list StrList) Union(t StrList) StrList {
	return append(list, t...)
}

func (list StrList) Intersection(t StrList) (r StrList) {
	minLen := int(math.Min(float64(len(list)), float64(len(t))))
	r = StrList(make([]string, minLen))
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

func (list StrList) Difference(t StrList) (r StrList) {
	minLen := int(math.Min(float64(len(list)), float64(len(t))))
	r = StrList(make([]string, minLen))
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
