func solveNQueens(n int) [][]string {
    board := make([][]byte, n)
    for r := 0; r < n; r++ {
        board[r] = make([]byte, n)
        for c := 0; c < n; c++ {
            board[r][c] = '.'
        }
    }
    res := [][]string{}
    dfs(0, board, &res)
    return res
}

func dfs(c int, board [][]byte, res *[][]string) {
    if c == len(board) {
        tmp := []string{}
        for r := 0; r < len(board); r++ {
            tmp = append(tmp, string(board[r][:]))
        }
        *res = append(*res, tmp)
        return
    }
    for r := 0; r < len(board); r++ {
        if isSafe(board, r, c) {
            board[r][c] = 'Q'
            dfs(c + 1, board, res)
            board[r][c] = '.'
        }
    }
}

func isSafe(board [][]byte, row, col int) bool {
    r, c := row, col
    for r >= 0 && c >= 0 {
        if board[r][c] == 'Q' { return false }
        r--
        c--
    }
    r, c = row, col
    for c >= 0 {
        if board[r][c] == 'Q' { return false }
        c--
    }
    r, c = row, col
    for c >= 0 && r < len(board) {
        if board[r][c] == 'Q' { return false }
        c--
        r++
    }
    return true
}
