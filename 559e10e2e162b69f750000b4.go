package kata

func Dominator(a []int) int {
	nums := make(map[int]int, len(a))

	for _, v := range a {
		nums[v] += 1
	}

	for ele, count := range nums {
		if count > len(a)/2 {
			return ele
		}
	}
	return -1
}
