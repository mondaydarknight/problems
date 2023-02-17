func diameterOfBinaryTree(root *TreeNode) int {
  maxf := 0
  dfs(root, &maxf)
  return maxf
}

func dfs(root *TreeNode, maxf *int) int {
  if root == nil { return 0 }
  left := dfs(root.Left, maxf)
  right := dfs(root.Right, maxf)
  *maxf = max(*maxf, left + right)
  return max(left, right) + 1
}

func max(a, b int) int {
  if a > b {
     return a
  }
  return b
}
