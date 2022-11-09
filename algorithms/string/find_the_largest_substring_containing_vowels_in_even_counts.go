func findTheLongestSubstring(s string) int {
    maxlen, curr, seen := 0, 0, make(map[int]int)
    seen[0] = -1
    for i := 0; i < len(s); i++ {
        curr ^= 1 << (strings.Index("aeiou", string(s[i])) + 1) >> 1
        if _, ok := seen[curr]; !ok {
            seen[curr] = i
        }
        maxlen = max(maxlen, i - seen[curr])
    }
    return maxlen
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
