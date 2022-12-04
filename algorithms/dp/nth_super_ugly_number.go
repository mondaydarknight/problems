func nthSuperUglyNumber(n int, primes []int) int {
    dp, dict := make([]int, n), make(map[int]int)
    dp[0] = 1
    for i := 0; i < len(primes); i++ {
        dict[primes[i]] = 0
    }
    for i := 1; i < n; i++ {
        dp[i] = math.MaxInt32
        for prime, t := range dict {
            dp[i] = min(dp[i], dp[t] * prime)
        }
        for prime, _ := range dict {
            if dp[i] % prime == 0 { dict[prime]++ }
        }
    }
    return dp[n - 1]
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
