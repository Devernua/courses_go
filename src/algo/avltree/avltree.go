package avltree

type node struct {
	left, right *node
	parent      *node
	rank        int // TODO: think about rank
	value       int // TODO: interface {}
}

type AVLTree struct {
	head *node
	// smth
}

func New() *AVLTree { return new(AVLTree) }

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

	// calc rank
	// TODO: if rank changed more then by one, need rotate, and complete calc rank
	rankNode := foundNode
	for rankNode != nil && getRank(rankNode.left) != getRank(rankNode.right) {
		rankNode.rank = max(getRank(rankNode.left), getRank(rankNode.right)) + 1
		rankNode = rankNode.parent
	}

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

	// TODO: if found head need unusual work
	parentNode := foundNode.parent

	// leaf, left == right == nil
	if foundNode.left == nil && foundNode.right == nil {
		if parentNode.left == foundNode {
			parentNode.left = nil
		} else {
			parentNode.right = nil
		}
		// if left || right == nil
	} else if foundNode.left == nil {
		if foundNode.value < parentNode.value {
			parentNode.left = foundNode.right
		} else {
			parentNode.right = foundNode.right
		}
	} else if foundNode.right == nil {
		if foundNode.value < parentNode.value {
			parentNode.left = foundNode.left
		} else {
			parentNode.right = foundNode.left
		}
		// find max, swap and delete leaf
	} else {
		maxFounded := foundNode.left
		for ; maxFounded.right != nil; maxFounded = maxFounded.right {
		}
		maxFounded.value, foundNode.value = foundNode.value, maxFounded.value
		if maxFounded.parent != foundNode {
			maxFounded.parent.right = nil // only if not left of head
		} else {
			maxFounded.parent.left = nil // only if eq left with founded node
		}
	}

	// calc rank
	// TODO: if rank changed more then by one, need rotate, and complete calc rank
	rankNode := parentNode
	for rankNode != nil && getRank(rankNode.left) != getRank(rankNode.right) {
		rankNode.rank = max(getRank(rankNode.left), getRank(rankNode.right)) + 1
		rankNode = rankNode.parent
	}

	// TODO: rotates
}

func getRank(n *node) int {
	if n == nil {
		return 0
	} else {
		return n.rank
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
