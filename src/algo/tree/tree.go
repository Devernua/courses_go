package tree

type ValueType int

type Node struct {
	Left, Right *Node
	Parent      *Node

	Value ValueType // key, some data
}

// TODO: My be not needed
type Tree struct {
	Head *Node
}

func (t *Tree) InOrder() []ValueType {
	result := make([]ValueType, 0)
	result = t.Head.inOrder(result)
	return result
}

func (n *Node) inOrder(accum []ValueType) []ValueType {
	if n.Left != nil {
		accum = n.Left.inOrder(accum)
	}
	accum = append(accum, n.Value)

	if n.Right != nil {
		accum = n.Right.inOrder(accum)
	}

	return accum
}

func (t *Tree) PreOrder() []ValueType {
	result := make([]ValueType, 0)
	result = t.Head.preOrder(result)
	return result
}

func (n *Node) preOrder(accum []ValueType) []ValueType {
	accum = append(accum, n.Value)

	if n.Left != nil {
		accum = n.Left.preOrder(accum)
	}

	if n.Right != nil {
		accum = n.Right.preOrder(accum)
	}

	return accum
}

func (t *Tree) PostOrder() []ValueType {
	result := make([]ValueType, 0)
	result = t.Head.postOrder(result)
	return result
}

func (n *Node) postOrder(accum []ValueType) []ValueType {
	if n.Left != nil {
		accum = n.Left.postOrder(accum)
	}

	if n.Right != nil {
		accum = n.Right.postOrder(accum)
	}

	accum = append(accum, n.Value)

	return accum
}
