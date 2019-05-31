package main

import (
	"algo/disjoinset"
	"fmt"
)

func main() {
	result := 1
	var n, e, d int
	fmt.Scanf("%v %v %v", &n, &e, &d)

	djset := disjoinset.New(n + 1)

	for i := 1; i <= n; i++ {
		djset.MakeSet(i)
	}

	for i := 0; i < e; i++ {
		var dst, src int
		fmt.Scanf("%v %v", &dst, &src)
		djset.Union(dst, src)
	}

	for i := 0; i < d; i++ {
		var dst, src int
		fmt.Scanf("%v %v", &dst, &src)
		if djset.Find(dst) == djset.Find(src) {
			result = 0
		}
	}

	fmt.Println(result)
}
