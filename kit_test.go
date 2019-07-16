package list_test

import (
	"testing"
	. "../list"
)

func TestSumInts(t *testing.T) {
	s := []int{1,1,1}
	total := SumInts(s)
	if total != 3 {
		t.Errorf("expect total of %v to be 3, but got %d", s, total)
	}
}