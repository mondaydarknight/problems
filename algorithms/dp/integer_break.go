func integerBreak(n int) int {
    if n == 2 { return 2 }
    if n == 3 { return 3 }
    dp := make([]int, n + 1)
    dp[0] = 1
    for i := 1; i <= n; i++ {
        for j := 1; j <= i; j++ {
            dp[i] = max(dp[i], dp[i - j] * j)
        }
    }
    return dp[len(dp) - 1]
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
