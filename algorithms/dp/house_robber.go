func rob(nums []int) int {
    memo := make(map[int]int)
    return max(f(nums, 0, memo), f(nums, 1, memo))
}

func f(nums []int, curr int, memo map[int]int) int {
    if curr >= len(nums) { return 0 }
    if sum, ok := memo[curr]; ok { return sum }
    memo[curr] = nums[curr] + max(f(nums, curr + 2, memo), f(nums, curr + 3, memo))
    return memo[curr]
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
