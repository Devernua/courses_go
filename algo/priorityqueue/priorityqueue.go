package priorityqueue


type PriorityQueue struct {
	arr []int
}

func (p *PriorityQueue) Push(val int) {}

func (p *PriorityQueue) Pop() (int, error) {
	return 0, nil
}
// TODO: method size if need

func (p PriorityQueue) parent(idx int) (parentIdx int) {
	parentIdx = idx / 2 // TODO: fix indexes
	return
}

func (p PriorityQueue) leftChild(idx int) (leftIdx int) {
	leftIdx = 2 * idx // TODO: fix idx
	return
}

func (p PriorityQueue) rightChild(idx int) (rightIdx int) {
	rightIdx = 2 * idx + 1
	return
}

func (p *PriorityQueue) shiftUp(idx int) {
	// TODO
}

func (p *PriorityQueue) shiftDown(idx int) {
	// TODO
}

