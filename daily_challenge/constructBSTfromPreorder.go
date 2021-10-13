package daily_challenge

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	i := 1
	for i < len(preorder) && preorder[i] < preorder[0] {
		i++
	}

	return &TreeNode{
		Val:   preorder[0],
		Left:  bstFromPreorder(preorder[1:i]),
		Right: bstFromPreorder(preorder[i:]),
	}
}
