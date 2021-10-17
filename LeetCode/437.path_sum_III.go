package main

import (
	"fmt"
)

func main() {
	result := pathSum(&TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode {
					Val: 3,
					Left: nil,
					Right: nil,
				},
				Right: &TreeNode {
					Val: -2,
					Left: nil,
					Right: nil,
				},
			},
			Right: &TreeNode{
				Val: 2,
				Left: nil,
				Right: &TreeNode {
					Val: 1,
					Left: nil,
					Right: nil,
				},
			},
		},
		Right: &TreeNode{
			Val: -3,
			Left: nil,
			Right: &TreeNode{
				Val: 11,
				Left: nil,
				Right: nil,
			},
		},
	}, 8)

	fmt.Println(result)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	var newRoot *TreeNodeWithParent = transferToNewStructure(root, nil)
	return numberOfSumPath(newRoot, targetSum)
}

func numberOfSumPath(root *TreeNodeWithParent, targetSum int) int {
	if root == nil {
		return 0
	}

	var count int

	if root.Left != nil {
		count += numberOfSumPath(root.Left, targetSum)
	}

	if root.Right != nil {
		count += numberOfSumPath(root.Right, targetSum)
	}

	sum := 0

	for current := root; current != nil; current = current.Parent {
		sum += current.Val

		if sum == targetSum {
			count++
		}
	}

	return count

}

type TreeNodeWithParent struct {
	Val    int
	Left   *TreeNodeWithParent
	Right  *TreeNodeWithParent
	Parent *TreeNodeWithParent
}

func transferToNewStructure(root *TreeNode, parent *TreeNodeWithParent) *TreeNodeWithParent {
	if root == nil {
		return nil
	}

	var newRoot *TreeNodeWithParent

	newRoot = &TreeNodeWithParent{
		Val:    root.Val,
		Left: nil,
		Right: nil,
		Parent: parent,
	}

	newRoot.Left = transferToNewStructure(root.Left, newRoot)
	newRoot.Right = transferToNewStructure(root.Right, newRoot)

	return newRoot
}
