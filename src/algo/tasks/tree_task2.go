package main

import (
	"fmt"
)

type Node2 struct {
	Left, Right *Node2
	Parent      *Node2

	Value   int // key, some data
	MaxNode *Node2
	MinNode *Node2
}

func main() {

	var numNodes int
	fmt.Scanf("%v", &numNodes)

	nodes := make([]Node2, numNodes)

	for i := 0; i < numNodes; i++ {
		var key, left, right int
		fmt.Scanf("%v %v %v", &key, &left, &right)
		curNode := &nodes[i]
		curNode.Value = key
		curNode.MaxNode = curNode
		curNode.MinNode = curNode
		if left != -1 {
			curNode.Left = &nodes[left]
			curNode.Left.Parent = curNode
			if curNode.Left.MaxNode != nil {
				if key <= curNode.Left.MaxNode.Value {
					fmt.Println("INCORRECT")
					return
				}
				curNode.MinNode = curNode.Left.MinNode
			}
		}

		if right != -1 {
			curNode.Right = &nodes[right]
			curNode.Right.Parent = curNode
			if curNode.Right.MaxNode != nil {
				if key >= curNode.Right.MinNode.Value {
					fmt.Println("INCORRECT")
					return
				}
				curNode.MaxNode = curNode.Right.MaxNode
			}
		}

		// TODO: down incorrect, need fix, need calc min and max and check full up
		if curNode.Parent != nil && curNode.Parent.Left == curNode {
			if curNode.Parent.Value <= curNode.MaxNode.Value {
				fmt.Println("INCORRECT")
				return
			}
			for ; curNode.Parent != nil && curNode.Parent.Left == curNode; curNode = curNode.Parent {
				curNode.Parent.MinNode = curNode.MinNode
			}

			if curNode.Parent != nil && curNode.MinNode.Value <= curNode.Parent.Value {
				fmt.Println("INCORRECT")
				return
			}
		}

		if curNode.Parent != nil && curNode.Parent.Right == curNode {
			if curNode.Parent.Value >= curNode.MinNode.Value {
				fmt.Println("INCORRECT")
				return
			}

			for ; curNode.Parent != nil && curNode.Parent.Right == curNode; curNode = curNode.Parent {
				curNode.Parent.MaxNode = curNode.MaxNode
			}

			if curNode.Parent != nil && curNode.MaxNode.Value >= curNode.Parent.Value {
				fmt.Println("INCORRECT")
				return
			}
		}
	}

	fmt.Println("CORRECT")
}
