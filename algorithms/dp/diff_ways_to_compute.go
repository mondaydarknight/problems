func diffWaysToCompute(expression string) []int {
    list := make([]int, 0)
    for i := 0; i < len(expression); i++ {
        if expression[i] == '+' || expression[i] == '-' || expression[i] == '*' {
            for _, s1 := range diffWaysToCompute(expression[:i]) {
                for _, s2 := range diffWaysToCompute(expression[i+1:]) {
                    c := 0
                    switch expression[i] {
                    case '+':
                        c = p1 + p2
                    case '-':
                        c = p1 - p2
                    case '*':
                        c = p1 * p2
                    }
                    list = append(list, c)
                }
            }
        }
    }
    if len(list) == 0 {
        n, _ := strconv.Atoi(expression)
        list = append(list, n)
    }
    return list
}
