package main

import (
	"fmt"
	"algo/priorityqueue"
)

type valType struct {
	id int
	time uint64
}

func cmp(l, r priorityqueue.ValueType) bool {
	if l.(valType).time == r.(valType).time {
		return l.(valType).id > r.(valType).id
	}
	return l.(valType).time > r.(valType).time
}

func main() {
	var curTime uint64 = 0
	var processorsNum, tasksNum int
	fmt.Scanf("%v %v", &processorsNum, &tasksNum)

	processors := priorityqueue.NewPriorityQueue(cmp)
	for i := 0; i < processorsNum; i++ {
		processors.Push(valType{i, 0})
	}

	for i := 0; i < tasksNum; i++ {
		var time uint64
		fmt.Scanf("%v", &time)

		proc := processors.Pop().(valType)
		curTime = proc.time
		fmt.Println(proc.id, curTime)
		proc.time += time
		processors.Push(proc)
	}
}

