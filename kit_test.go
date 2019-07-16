package list_test

import (
	"testing"
	. "../list"
)

func TestSumInts(t *testing.T) {
	s := []int{1,1,1,1,1,1}
	total := SumInts(s)
	if total != 6 {
		t.Errorf("expect total of %v to be 6, but got %d", s, total)
	}
}

func TestSumFloat64(t *testing.T) {
	s := []float64{1.5,1.5,1.5,1.5,1.5,1.5}
	total := SumFloat64s(s)
	if total != 9 {
		t.Errorf("expect total of %v to be 9, but got %.2f", s, total)
	}
}