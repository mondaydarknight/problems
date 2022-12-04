func maxCoins(nums []int) int {
    n, inums := 1, make([]int, len(nums) + 2)    
    for _, num := range nums {
        if num > 0 { inums[n], n = num, n + 1 }
    }
    inums[0], inums[n], n = 1, 1, n + 1
    dp := make([][]int, n)
    for i := 0; i < len(dp); i++ {
        dp[i] = make([]int, n)
    }
    for k := 2; k < n; k++ {
        for left := 0; left < n - k; left++ {
            right := left + k
            for i := left + 1; i < right; i++ {
                dp[left][right] = max(dp[left][right], 
                inums[left] * inums[i] * inums[right] + dp[left][i] + dp[i][right])
            }
        }
    }
    return dp[0][n - 1]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
