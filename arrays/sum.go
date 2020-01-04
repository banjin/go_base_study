package main

func Sum(a []int) int{
	n := 0
	for _, value := range(a) {
		n += value
	}
	return n
}

func SumAll(a ...[]int) (sums []int) {
	lengthOfNumbers := len(a)
	sums = make([]int, lengthOfNumbers)

	for i, numbers := range a {
        sums[i] = Sum(numbers)  // 使用切片的索引访问切片内的元素, 使用 = 对切片元素进行赋值。
	}
	return 
}