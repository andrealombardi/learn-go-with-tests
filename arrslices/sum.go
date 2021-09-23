package arrslices

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(nums ...[]int) []int {

	var sums []int
	for _, num := range nums {
		sums = append(sums, Sum(num))
	}
	return sums
}

func SumAllTails(nums ...[]int) []int {

	var sums []int
	for _, num := range nums {
		if len(num) == 0 {
			sums = append(sums, 0)
		} else {
			tail := num[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums

}
