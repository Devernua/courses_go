package main

import "fmt"
import "algo/priorityqueue"

func main() {
	q := priorityqueue.PriorityQueue{}

	
	var n int
	fmt.Scanf("%d", &n)

	fmt.Println("Hello", q.Size())
}
