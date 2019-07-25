package list

type IntList []int

type ILEachFn func(v, i int)

func (l IntList) Each(f ILEachFn) IntList {
	for i, v := range l {
		f(v, i)
	}
	return l
}
