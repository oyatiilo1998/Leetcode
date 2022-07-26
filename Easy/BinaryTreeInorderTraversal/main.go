package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	var tree *TreeNode = &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}

	fmt.Println(inorderTraversal(tree))
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	if root.Left != nil && root.Right != nil {
		return append(append(inorderTraversal(root.Left), root.Val), inorderTraversal(root.Right)...)
	}
	if root.Left != nil && root.Right == nil {
		return append(inorderTraversal(root.Left), root.Val)
	}
	if root.Left == nil && root.Right != nil {
		return append([]int{root.Val}, inorderTraversal(root.Right)...)
	}
	return []int{}
}
