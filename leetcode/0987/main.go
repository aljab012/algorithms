package main

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type NodeWithInfo struct {
	val, row, col int
}

func verticalTraversal(root *TreeNode) [][]int {
	nodes := []*NodeWithInfo{}

	var visit func(node *TreeNode, row, col int)
	visit = func(node *TreeNode, row, col int) {
		if node == nil {
			return
		}
		nodes = append(nodes, &NodeWithInfo{val: node.Val, row: row, col: col})
		visit(node.Left, row+1, col-1)
		visit(node.Right, row+1, col+1)
	}
	visit(root, 0, 0)

	// sort by col, row
	sort.Slice(nodes, func(i, j int) bool {
		if nodes[i].col != nodes[j].col {
			return nodes[i].col < nodes[j].col
		}
		if nodes[i].row != nodes[j].row {
			return nodes[i].row <= nodes[j].row
		}
		return nodes[i].val <= nodes[j].val
	})

	result := [][]int{}
	curCol := nodes[0].col
	colValues := []int{}
	for _, node := range nodes {
		if node.col == curCol {
			colValues = append(colValues, node.val)
		} else {
			result = append(result, colValues)
			colValues = []int{node.val}
			curCol = node.col
		}
	}
	result = append(result, colValues)
	return result
}
