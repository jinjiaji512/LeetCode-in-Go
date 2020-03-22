package problem0102

import (
	"fmt"
	"github.com/jinjiaji512/LeetCode-in-Go/kit"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// inOrderTraversal
func levelOrder(root *kit.TreeNode) [][]int {
	result := [][]int{}
	result = inOrderTraversal(root, 1, result)
	return result
}

// @title	inOrderTraversal
// @description	前序遍历的方式递归调用
// @auth	jinjiaji 2020年03月21日
// @param	node	*kit.TreeNode	tree节点指针
// @param	level	int	tree节点所处的层
// @param	result	[][]int	层序遍历结果存储
// @return
func inOrderTraversal(node *kit.TreeNode, level int, result [][]int) [][]int {
	fmt.Println(cap(result), len(result))
	if node == nil {
		return result
	}
	//初始化层级数组
	if len(result) < level {
		result = append(result, []int{})
	}
	//前序遍历
	if node.Left != nil {
		result = inOrderTraversal(node.Left, level+1, result)
	}
	result[level-1] = append(result[level-1], node.Val)
	if node.Right != nil {
		result = inOrderTraversal(node.Right, level+1, result)
	}
	return result
}

// levelOrder2 闭包的方式传递result
func levelOrder2(root *kit.TreeNode) [][]int {
	result := [][]int{}
	var dfs func(*kit.TreeNode, int)
	dfs = func(node *kit.TreeNode, level int) {
		if node == nil {
			return
		}
		//初始化层级数组
		if len(result) < level {
			result = append(result, []int{})
		}
		//前序遍历
		if node.Left != nil {
			dfs(node.Left, level+1)
		}
		result[level-1] = append(result[level-1], node.Val)
		if node.Right != nil {
			dfs(node.Right, level+1)
		}
	}
	dfs(root, 1)
	return result
}
