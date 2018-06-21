package avltree

type node struct {
	left, right *node
	parent *node
	rank int // TODO: think about rank
	value int // TODO: interface {}
}

type AVLTree struct {
	head *node
	// smth
}

func (t *AVLTree) Insert(value int) error {
	if t.head == nil {
		t.head = &node{rank: 1, value: value}
		return nil
	}

	foundNode := t.head
	var newNode *node

	for newNode == nil {
		if value < foundNode.value {
			if foundNode.left != nil {
				foundNode = foundNode.left
			} else {
				newNode = &node{rank: 1, value: value, parent: foundNode}
				foundNode.left = newNode
			}
		} else { // >=
			if foundNode.right != nil {
				foundNode = foundNode.right
			} else {
				newNode = &node{rank: 1, value: value, parent: foundNode}
				foundNode.right = newNode
			}
		}
	}

	// TODO: calc ranks
	// TODO: rotates
	return nil
}

func (t *AVLTree) Delete(value int) {
	foundNode := t.head
	for foundNode != nil && foundNode.value != value {
		if value < foundNode.value {
			foundNode = foundNode.left
		} else {
			foundNode = foundNode.right
		}
	}

	if foundNode == nil {
		return
	}

	if foundNode.left == nil && foundNode.right == nil {
		p := foundNode.parent
		if foundNode.value < p.value {
			p.left = nil
		} else {
			p.right = nil
		}
	} else if foundNode.left == nil {
		p := foundNode.parent
		if foundNode.value < p.value {
			p.left = foundNode.right
		} else {
			p.right = foundNode.right
		}
	} else if foundNode.right == nil {

	}

	// TODO: calc ranks
	// TODO: rotates
}