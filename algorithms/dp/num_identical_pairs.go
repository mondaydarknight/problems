func numIdenticalPairs(nums []int) int {
	m := make(map[int]int)
	res := 0
	for _, n := range nums {
		m[n]++
	}
	for _, count := range m {
		if count > 1 {
			sum := 0
			for i := count - 1; i >= 1; i-- {
				sum += i
			}
			res += sum
		}
	}
	return res
}
