package main

import (
	"algo/tree"
	"fmt"
)

func main() {

	var numNodes int
	fmt.Scanf("%v", &numNodes)

	nodes := make([]tree.Node, numNodes)

	for i := 0; i < numNodes; i++ {
		var key, left, right int
		fmt.Scanf("%v %v %v", &key, &left, &right)
		curNode := &nodes[i]
		curNode.Value = tree.ValueType(key)
		if left != -1 {
			curNode.Left = &nodes[left]
			curNode.Left.Parent = curNode
		}

		if right != -1 {
			curNode.Right = &nodes[right]
			curNode.Right.Parent = curNode
		}
	}

	head := &nodes[0]
	for ; head.Parent != nil; head = head.Parent {
	}
	t := tree.Tree{Head: head}

	for _, i := range t.InOrder() {
		fmt.Printf("%v ", i)
	}
	fmt.Println()

	for _, i := range t.PreOrder() {
		fmt.Printf("%v ", i)
	}
	fmt.Println()

	for _, i := range t.PostOrder() {
		fmt.Printf("%v ", i)
	}
	fmt.Println()
}
