package main


func Sum(a [5]int) int{
	n := 0
	for i := 0; i < len(a); i++ {
		n += a[i]
	}
	return n
}

