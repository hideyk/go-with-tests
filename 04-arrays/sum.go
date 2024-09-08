package main

func Sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func SumAll(arrayOfNums ...[]int) []int {
	var sums []int

	for _, nums := range arrayOfNums {
		sums = append(sums, Sum(nums))
	}

	return sums
}
