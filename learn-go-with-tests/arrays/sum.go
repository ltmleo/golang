package arrays

func Sum(numbers []int) int {
	res := 0
	for _, n := range numbers {
		res += n
	}
	return res
}

func SumAll(numbers ... []int ) []int {
	var res []int
    for _, n := range numbers {
        res = append(res, Sum(n))
    }
	return res
}