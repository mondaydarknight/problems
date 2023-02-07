func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	res := math.MaxInt32
	sums := make([]int, n)
	for i := 1; i < n; i++ {
		sums[i] = sums[i-1] + nums[i - 1]
	}
	for i := 0; i < n; i++ {
		l, r := 0, n
		for l < r {
			mid := (l + r) / 2
			sum := sums[mid] - sums[i] + nums[mid]
			if sum >= target {
				res = min(res, mid-i+1)
				r = mid
			} else {
				l = mid + 1
			}
		}
	}
	if res == math.MaxInt32 {
		return 0
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
