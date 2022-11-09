func maxIncreaseKeepingSkyline(grid [][]int) int {
    rowMaxes, colMaxes := make([]int, len(grid)), make([]int, len(grid))
    for r := 0; r < len(grid); r++ {
        for c := 0; c < len(grid[r]); c++ {
            rowMaxes[r] = max(rowMaxes[r], grid[r][c])
            colMaxes[c] = max(colMaxes[c], grid[r][c])
        }
    }
    res := 0
    for r := 0; r < len(grid); r++ {
        for c := 0;  c < len(grid); c++ {
            res += min(rowMaxes[r], colMaxes[c]) - grid[r][c]
        }
    }
    return res
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
