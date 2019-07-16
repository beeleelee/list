package list 

func SumInts(l []int) int {
	total := 0 
	for _, v := range l {
		total += v
	}
	return total
}

func SumFloat64s(l []float64) float64 {
	total := 0.0
	for _, v := range l {
		total += v
	}
	return total
}

