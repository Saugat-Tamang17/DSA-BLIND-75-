package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil || t == nil {
		return false
	}
	if s.Val != t.Val {
		return false
	}
	return isSameTree(s.Left, t.Left) && isSameTree(s.Right, t.Right)
}
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	if isSameTree(root, subRoot) {
		return true
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}
func main() {
	// Main tree
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 5}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 2}
	// Subtree
	subRoot := &TreeNode{Val: 4}
	subRoot.Left = &TreeNode{Val: 1}
	subRoot.Right = &TreeNode{Val: 2}
	result := isSubtree(root, subRoot)
	fmt.Println("Is subtree:", result)
}
