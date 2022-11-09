func lengthOfLongestSubstring(s string) int {
    pivot, length, hash := 0, 0, make(map[byte]int)
    for i := 0; i < len(s); i++ {
        if v, ok := hash[s[i]]; ok {
            pivot = max(pivot, v + 1)
        }
        hash[s[i]] = i
        length = max(length, i - pivot + 1)
    }
    return length
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
