package util

func SumSlice(s []int) int {
	sum := 0
	for _, n := range s {
		sum += n
	}
	return sum
}

func MulSlice(s []int) int {
	product := 1
	for _, n := range s {
		product *= n
	}
	return product
}

func MinSlice(s []int) int {
	min := s[0]
	for _, n := range s {
		if n < min {
			min = n
		}
	}
	return min
}

func MaxSlice(s []int) int {
	max := s[0]
	for _, n := range s {
		if n > max {
			max = n
		}
	}
	return max
}
