func numDecodings(s string) int {
    dp := make([]int, len(s) + 1)
    dp[0] = 1
    dp[1] = 1
    if s[0] == '0' { dp[1] = 0 }
    for i := 2; i <= len(s); i++ {
        d1, d2 := s[i-1]-'0', 10*(s[i-2]-'0') + s[i-1]-'0'
        if d1 >= 1 { dp[i] += dp[i - 1] }
        if d2 >= 10 && d2 <= 26 { dp[i] += dp[i - 2] }
    }
    return dp[len(s)]
}
