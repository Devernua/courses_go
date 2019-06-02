package main

import (
	"fmt"
)

type Node struct {
	Left, Right *Node
	Parent      *Node

	Value int // key, some data

	Visited bool
}

// TODO: My be not needed
type Tree struct {
	Head *Node
}

func (t *Tree) NewInOrder() {
	curNode := t.Head

	for !t.Head.Visited {
		if curNode.Left != nil && !curNode.Left.Visited {
			curNode = curNode.Left
		} else if curNode.Right != nil && !curNode.Right.Visited {
			fmt.Printf("%v ", curNode.Value)
			curNode = curNode.Right
		} else {
			if curNode.Right == nil {
				fmt.Printf("%v ", curNode.Value)
			}
			curNode.Visited = true
			curNode = curNode.Parent
		}
	}
}

func (t *Tree) NewPreOrder() {
	curNode := t.Head

	for t.Head.Visited {
		if curNode.Left != nil && curNode.Left.Visited {
			fmt.Printf("%v ", curNode.Value)
			curNode = curNode.Left
		} else if curNode.Right != nil && curNode.Right.Visited {
			if curNode.Left == nil {
				fmt.Printf("%v ", curNode.Value)
			}
			curNode = curNode.Right
		} else {
			if curNode.Left == nil && curNode.Right == nil {
				fmt.Printf("%v ", curNode.Value)
			}
			curNode.Visited = false
			curNode = curNode.Parent
		}
	}
}

func (t *Tree) NewPostOrder() {
	curNode := t.Head

	for !t.Head.Visited {
		if curNode.Left != nil && !curNode.Left.Visited {
			curNode = curNode.Left
		} else if curNode.Right != nil && !curNode.Right.Visited {
			curNode = curNode.Right
		} else {
			fmt.Printf("%v ", curNode.Value)
			curNode.Visited = true
			curNode = curNode.Parent
		}
	}
}

func (t *Tree) InOrder() []int {
	result := make([]int, 0)
	result = t.Head.inOrder(result)
	return result
}

func (n *Node) inOrder(accum []int) []int {
	if n.Left != nil {
		accum = n.Left.inOrder(accum)
	}
	accum = append(accum, n.Value)

	if n.Right != nil {
		accum = n.Right.inOrder(accum)
	}

	return accum
}

func (t *Tree) PreOrder() []int {
	result := make([]int, 0)
	result = t.Head.preOrder(result)
	return result
}

func (n *Node) preOrder(accum []int) []int {
	accum = append(accum, n.Value)

	if n.Left != nil {
		accum = n.Left.preOrder(accum)
	}

	if n.Right != nil {
		accum = n.Right.preOrder(accum)
	}

	return accum
}

func (t *Tree) PostOrder() []int {
	result := make([]int, 0)
	result = t.Head.postOrder(result)
	return result
}

func (n *Node) postOrder(accum []int) []int {
	if n.Left != nil {
		accum = n.Left.postOrder(accum)
	}

	if n.Right != nil {
		accum = n.Right.postOrder(accum)
	}

	accum = append(accum, n.Value)

	return accum
}

func main() {

	var numNodes int
	fmt.Scanf("%v", &numNodes)

	nodes := make([]Node, numNodes)

	for i := 0; i < numNodes; i++ {
		var key, left, right int
		fmt.Scanf("%v %v %v", &key, &left, &right)
		curNode := &nodes[i]
		curNode.Value = key
		curNode.Visited = false
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
	t := Tree{Head: head}

	t.NewInOrder()
	fmt.Println()

	t.NewPreOrder()
	fmt.Println()

	t.NewPostOrder()
	fmt.Println()
}

//
//10
//0 7 2
//10 -1 -1
//20 -1 6
//30 8 9
//40 3 -1
//50 -1 -1
//60 1 -1
//70 5 4
//80 -1 -1
//90 -1 -1
