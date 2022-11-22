func wordBreak(s string, wordDict []string) []string {
    list := []string{}
    dict := make(map[string]bool)
    for _, word := range wordDict {
        dict[word] = false
    }
    dfs(s, "", &list, dict, 0)
    return list
}

func dfs(s, ss string, list *[]string, dict map[string]bool, c int) bool {
    if c == len(s) { return true }
    for i := c; i < len(s); i++ {
        if _, ok := dict[s[c:i+1]]; ok {
            dict[s[c:i+1]] = true
            if dfs(s, ss + s[c:i+1] + " ", list, dict, i + 1) {
                *list = append(*list, ss + s[c:i+1])
            }
            dict[s[c:i+1]] = false
        }
    }
    return false
}
