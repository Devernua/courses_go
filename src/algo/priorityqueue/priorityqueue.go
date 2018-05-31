package priorityqueue

type KeyType interface{} // priority
type ValueType interface{}

type item struct {
	priority KeyType
	value ValueType
}

type PriorityQueue struct {
	arr []item
	comparator func(KeyType, KeyType) bool
}

func (p *PriorityQueue) Push(key KeyType, val ValueType) {
	p.arr = append(p.arr, item{key, val})
	p.shiftUp(len(p.arr) - 1)
}

func (p *PriorityQueue) Pop() (KeyType, ValueType) {
	if len(p.arr) == 0 {
		panic("queue is empty")
	}
	
	result := p.arr[0]
	p.arr[0] = p.arr[len(p.arr) - 1]
	p.arr = p.arr[:len(p.arr) - 1]
	p.shiftDown(0)
	return result, nil
}

func (p PriorityQueue) Size() (int) {
	return len(p.arr)
}

func (p PriorityQueue) parent(idx int) (parentIdx int) {
	if idx == 0 {
		panic("Failed to get parent from root")
	}

	parentIdx = ((idx + 1) / 2) - 1
	return
}

func (p PriorityQueue) leftChild(idx int) (leftIdx int) {
	leftIdx = 2 * idx + 1
	return
}

func (p PriorityQueue) rightChild(idx int) (rightIdx int) {
	rightIdx = 2 * idx + 2
	return
}

func (p *PriorityQueue) shiftUp(idx int) {
	for idx < len(p.arr) && idx > 0 && p.comparator(p.arr[p.parent(idx)].priority, p.arr[idx].priority) {
		p.arr[p.parent(idx)], p.arr[idx] = p.arr[idx], p.arr[p.parent(idx)]
		idx = p.parent(idx)
	}
}

func (p *PriorityQueue) shiftDown(idx int) {
	maxIdx := idx

	leftIdx := p.leftChild(idx)
	if leftIdx < len(p.arr) && p.comparator(p.arr[maxIdx].priority, p.arr[leftIdx].priority) {
		maxIdx = leftIdx
	}

	rightIdx := p.rightChild(idx)
	if rightIdx < len(p.arr) && p.comparator(p.arr[maxIdx].priority, p.arr[rightIdx]) {
		maxIdx = rightIdx
	}

	if idx != maxIdx {
		p.arr[idx], p.arr[maxIdx] = p.arr[maxIdx], p.arr[idx]
		p.shiftDown(maxIdx)
	}
}
