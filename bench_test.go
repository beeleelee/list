package list_test

// import (
// 	"testing"
// 	"time"

// 	. "github.com/beeleelee/list"
// )

// func TestBigList(t *testing.T) {
// 	maxSize := 1 << 26
// 	t.Log(maxSize)
// 	list := List(make([]Item, maxSize))
// 	intSlice := make([]int, maxSize)
// 	for i, _ := range list {
// 		list[i] = i
// 		intSlice[i] = i
// 	}
// 	listSumStartTime := time.Now()
// 	s := list.Reduce(func(a, b Item) Item {
// 		return a.(int) + b.(int)
// 	}, nil)
// 	listSumEndTime := time.Now()
// 	t.Logf("list sum duration: %v", listSumEndTime.Sub(listSumStartTime))

// 	sliceSumStartTime := time.Now()
// 	sum := 0
// 	for _, v := range intSlice {
// 		sum += v
// 	}
// 	sliceSumEndTime := time.Now()
// 	t.Logf("slice sum duration: %v", sliceSumEndTime.Sub(sliceSumStartTime))
// 	t.Log(s, sum)
// 	t.Fail()
// }
// func concurrentSum(c chan int, l List, times int) {
// 	size := len(l)
// 	splitNum := size / times
// 	var start, end, count int
// 	count = 0
// 	for i := 0; i < times; i++ {
// 		start = i * splitNum
// 		if i+1 == times {
// 			end = size
// 		} else {
// 			end = (i + 1) * splitNum
// 		}
// 		go func(l List) {
// 			count++
// 			c <- (l.Reduce(func(a, b Item) Item {
// 				return a.(int) + b.(int)
// 			}, nil)).(int)
// 			if count+1 == times {
// 				close(c)
// 			}
// 		}(l[start:end])
// 	}
// }
// func TestBigListConcurrent(t *testing.T) {
// 	maxSize := 1 << 26
// 	t.Log(maxSize)
// 	list := List(make([]Item, maxSize))
// 	intSlice := make([]int, maxSize)
// 	for i, _ := range list {
// 		list[i] = i
// 		intSlice[i] = i
// 	}
// 	listSumStartTime := time.Now()
// 	goRoutineNum := 10
// 	s := 0

// 	c := make(chan int)
// 	concurrentSum(c, list, goRoutineNum)
// 	for i := 0; i < goRoutineNum; i++ {
// 		s += <-c
// 	}
// 	listSumEndTime := time.Now()
// 	t.Logf("list sum duration: %v", listSumEndTime.Sub(listSumStartTime))

// 	sliceSumStartTime := time.Now()
// 	sum := 0
// 	for _, v := range intSlice {
// 		sum += v
// 	}
// 	sliceSumEndTime := time.Now()
// 	t.Logf("slice sum duration: %v", sliceSumEndTime.Sub(sliceSumStartTime))
// 	t.Log(s, sum)
// 	t.Fail()
// }
