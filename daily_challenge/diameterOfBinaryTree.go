package daily_challenge

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	res := 0
	dfs(root, &res)
	return res
}

func dfs(node *TreeNode, res *int) int {
	if node == nil {
		return 0
	}
	left := dfs(node.Left, res)
	right := dfs(node.Right, res)
	*res = max(*res, left+right)
	return max(left, right) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
