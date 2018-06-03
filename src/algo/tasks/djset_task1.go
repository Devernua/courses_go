package main

import (
	"algo/disjoinset"
	"fmt"
)

func main() {
	var n, m int
	fmt.Scanf("%v %v", &n, &m)

	weights := make(map[int]int, n)
	djset := disjoinset.NewDisJoinSet()
	maxWeight := 0

	for i := 1; i <= n; i++ {
		var w int
		fmt.Scanf("%v", &w)
		weights[i] = w
		djset.MakeSet(i)
		if w > maxWeight {
			maxWeight = w
		}
	}

	for i := 0; i < m; i++ {
		var dst, src int
		fmt.Scanf("%v %v", &dst, &src)
		l := djset.Find(dst)
		r := djset.Find(src)

		if l != r {
			weights[l], weights[r] = weights[l] + weights[r], weights[l] + weights[r]
			djset.Union(l, r)
			if weights[l] > maxWeight {
				maxWeight = weights[l]
			}
		}

		fmt.Println(maxWeight)
	}
}