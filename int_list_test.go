package list_test

import (
	"testing"
	//"reflect"
	. "../list"
)

func TestIntListEach(t *testing.T) {
	list := IntList([]int{2,3,5,7,9,11,13})
	list.Each(func(v int, i int){
		t.Log(v, i)
	})
	t.Fail()
}