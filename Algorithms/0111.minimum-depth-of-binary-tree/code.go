package problem0111

import "github.com/jinjiaji512/LeetCode-in-Go/kit"

func minDepth(root *kit.TreeNode) int {
	//递归实现,左右分别递归
	//if root == nil return 0
	//else depth ++
	//depth += min(depth(left),depth(right))

	depth := 0
	if root == nil {
		//节点为nil，深度为0
		return depth
	}
	//不为nil，深度+1
	depth++
	//子树深度
	depth += minAndNotZero(minDepth(root.Left), minDepth(root.Right))
	return depth
}

func minAndNotZero(a, b int) int {
	if a == 0 {
		return b
	} else if b == 0 {
		return a
	} else if a > b {
		return b
	} else {
		return a
	}
}

//递归 异步剪枝操作
func minDepth2(root *kit.TreeNode) int {
	//递归实现,左右分别递归
	//if root == nil return 0
	//else depth ++
	//depth += min(depth(left),depth(right))

	depth := 0
	if root == nil {
		//节点为nil，深度为0
		return depth
	}
	//不为nil，深度+1
	depth++
	//子树深度
	depth += minAndNotZero(minDepth(root.Left), minDepth(root.Right))
	return depth
}
