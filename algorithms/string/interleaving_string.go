func isInterleave(s1 string, s2 string, s3 string) bool {
    if len(s1) + len(s2) != len(s3) { return false }
    memo := make([][]int, len(s1))
    for i := range memo {
        memo[i] = make([]int, len(s2))
        for j := range memo[i] {
            memo[i][j] = -1
        }
    }
    return interleave(s1, s2, s3, 0, 0, 0, memo)
}

func interleave(s1, s2, s3 string, m, n, k int, memo [][]int) bool {
    if m == len(s1) { return s2[n:] == s3[k:] }
    if n == len(s2) { return s1[m:] == s3[k:] }
    if memo[m][n] >= 0 { return memo[m][n] == 1 }
    memo[m][n] = 0
    if s1[m] == s3[k] && interleave(s1, s2, s3, m + 1, n, k + 1, memo) ||
        s2[n] == s3[k] && interleave(s1, s2, s3, m, n + 1, k + 1, memo) {
        memo[m][n] = 1
    }
    return memo[m][n] == 1
}
