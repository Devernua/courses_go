package main

import "fmt"

func siftDownWithPrint(arr *[]int, idx int, results *[][2]int) {
	maxIdx := idx

	leftIdx := 2*idx + 1
	if leftIdx < len(*arr) && (*arr)[leftIdx] < (*arr)[maxIdx] {
		maxIdx = leftIdx
	}

	rightIdx := 2*idx + 2
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
