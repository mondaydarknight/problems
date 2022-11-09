func pathSum(root *TreeNode, targetSum int) (paths [][]int) {
    bfs(root, targetSum, []int{}, &paths)
    return
}

func bfs(root *TreeNode, target int, path []int, paths *[][]int) {
    if root == nil { return }
    target -= root.Val
    path = append(path, root.Val)
    if root.Left == nil && root.Right == nil && target == 0 {
        *paths = append(*paths, append([]int{}, path...))
        return
    }
    bfs(root.Left, target, path, paths)
    bfs(root.Right, target, path, paths)
}
