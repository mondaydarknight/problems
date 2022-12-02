func removeDuplicateLetters(s string) string {
    lastOcc, stack, visited := make(map[byte]int), []byte{}, make(map[byte]bool)
    for i := 0; i < len(s); i++ {
        lastOcc[s[i]] = i
    }
    for i := 0; i < len(s); i++ {
        if _, ok := visited[s[i]]; !ok {
            for len(stack) > 0 && stack[len(stack)-1] > s[i] && lastOcc[stack[len(stack)-1]] > i {
                delete(visited, stack[len(stack)-1])
                stack = stack[:len(stack)-1]
            }
            stack = append(stack, s[i])
            visited[s[i]] = true
        }
    }
    return string(stack)
}
