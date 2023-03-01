func maxAreaOfIsland(grid [][]int) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				res = max(res, dfs(grid, i, j)+1)
			}
		}
	}
	return res
}

func dfs(grid [][]int, i, j int) int {
	count := 0
	grid[i][j] = 0
	if i > 0 && grid[i-1][j] == 1 {
		count += dfs(grid, i-1, j) + 1
	}
	if i < len(grid)-1 && grid[i+1][j] == 1 {
		count += dfs(grid, i+1, j) + 1
	}
	if j > 0 && grid[i][j-1] == 1 {
		count += dfs(grid, i, j-1) + 1
	}
	if j < len(grid[i])-1 && grid[i][j+1] == 1 {
		count += dfs(grid, i, j+1) + 1
	}
	return count
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
