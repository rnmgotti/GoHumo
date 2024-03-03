package taskd7

func Split(arr []int) (even []int, odd []int) {
	for _, v := range arr {
		if v%2 == 0 {
			even = append(even, v)
		} else {
			odd = append(odd, v)
		}
	}
	return even, odd
}
