package main

import "fmt"


type PriorityQueue struct {
	arr []int
}

func (p *PriorityQueue) Push(val int) {
	p.arr = append(p.arr, val)
	p.shiftUp(len(p.arr) - 1)
}

func (p *PriorityQueue) Pop() (int, error) {
	if len(p.arr) == 0 {
		return 0, fmt.Errorf("queue is empty")
	}
	result := p.arr[0]
	p.arr[0] = p.arr[len(p.arr) - 1]
	p.arr = p.arr[:len(p.arr) - 1]
	p.shiftDown(0)
	return result, nil
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
	for idx < len(p.arr) && idx > 0 && p.arr[p.parent(idx)] < p.arr[idx] {
		p.arr[p.parent(idx)], p.arr[idx] = p.arr[idx], p.arr[p.parent(idx)]
		idx = p.parent(idx)
	}
}

func (p *PriorityQueue) shiftDown(idx int) {
	maxIdx := idx

	leftIdx := p.leftChild(idx)
	if leftIdx < len(p.arr) && p.arr[leftIdx] > p.arr[maxIdx] {
		maxIdx = leftIdx
	}

	rightIdx := p.rightChild(idx)
	if rightIdx < len(p.arr) && p.arr[rightIdx] > p.arr[maxIdx] {
		maxIdx = rightIdx
	}

	if idx != maxIdx {
		p.arr[idx], p.arr[maxIdx] = p.arr[maxIdx], p.arr[idx]
		p.shiftDown(maxIdx)
	}
}

func siftDownWithPrint(arr *[]int, idx int, results *[][2]int) {
	maxIdx := idx

	leftIdx := 2 * idx + 1
	if leftIdx < len(*arr) && (*arr)[leftIdx] < (*arr)[maxIdx] {
		maxIdx = leftIdx
	}

	rightIdx := 2 * idx + 2
	if rightIdx < len(*arr) && (*arr)[rightIdx] < (*arr)[maxIdx] {
		maxIdx = rightIdx
	}

	if idx != maxIdx {
		(*arr)[idx], (*arr)[maxIdx] = (*arr)[maxIdx], (*arr)[idx]
		*results = append(*results, [2]int{idx, maxIdx})
		siftDownWithPrint(arr, maxIdx, results)
	}
}

func main() {
	var n int
	fmt.Scanf("%d", &n)

	arr := make([]int, n)
	for i := range arr {
		fmt.Scanf("%d", &arr[i])
	}

	var results [][2]int
	for i := (len(arr) - 1) / 2; i != 0; i-- {
		siftDownWithPrint(&arr, i, &results)
	}
	siftDownWithPrint(&arr, 0, &results)

	fmt.Println(len(results))
	for _, pair := range results {
		fmt.Println(pair[0], pair[1])
	}
}
