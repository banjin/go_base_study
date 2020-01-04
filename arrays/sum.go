package main

func Sum(a []int) int{
	n := 0
	for _, value := range(a) {
		n += value
	}
	return n
}

