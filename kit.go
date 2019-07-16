package list 

func SumInts(l []int) int {
	total := 0 
	for _, v := range l {
		total += v
	}
	return total
}