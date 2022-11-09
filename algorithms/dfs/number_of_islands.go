func numIslands(grid [][]byte) int {
    m, n, res := len(grid), len(grid[0]), 0
    for row := 0; row < m; row++ {
        if grid[row][0] == '1' {
            dfs(grid, row, 0, m, n)
            res++
        }
        if grid[row][n - 1] == '1' {
            dfs(grid, row, n - 1, m, n)
            res++
        }
    }
    for col := 0; col < n; col++ {
        if grid[0][col] == '1' {
            dfs(grid, 0, col, m, n)
            res++
        }
        if grid[m - 1][col] == '1' {
            dfs(grid, m - 1, col, m, n)
            res++
        }
    }
    for row := 0; row < m; row++ {
        for col := 0; col < n; col++ {
            if grid[row][col] == '1' {
                dfs(grid, row, col, m, n)
                res++
            }
        }
    }
    return res
}

func dfs(grid [][]byte, row, col, m, n int) {
    grid[row][col] = '2'
    if row > 0 && grid[row - 1][col] == '1' {
        dfs(grid, row - 1, col, m, n)
    }
    if row < m - 1 && grid[row + 1][col] == '1' {
        dfs(grid, row + 1, col, m, n)
    }
    if col > 0 && grid[row][col - 1] == '1' {
        dfs(grid, row, col - 1, m, n)
    }
    if col < n - 1 && grid[row][col + 1] == '1' {
        dfs(grid, row, col + 1, m, n)
    }
}
