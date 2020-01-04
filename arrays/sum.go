package main

func Sum(a []int) int{
	n := 0
	for _, value := range(a) {
		n += value
	}
	return n
}

func SumAll(a ...[]int) []int {
	// lengthOfNumbers := len(a)
	// sums = make([]int, lengthOfNumbers)

	var sums []int
	for _, numbers := range a {
		// sums[i] = Sum(numbers)  // 使用切片的索引访问切片内的元素, 使用 = 对切片元素进行赋值。
		sums = append(sums, Sum(numbers))
	
	}
	return sums
}

// 去除第一个元素的剩余元素的和
func SumAllTail(a ...[]int) []int{
	var sums []int
	for _,  numbers  := range a {
		numbers = numbers[1:]
		sums = append(sums, Sum(numbers))
	}
	return sums

}