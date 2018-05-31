package priorityqueue

import (
	"testing"
)

func cmpLessInt(l, r KeyType) bool {
	return l.(int) < r.(int)
}

func cmpMoreInt(l, r KeyType) bool {
	return r.(int) < l.(int)
}

func TestSortInt(t *testing.T) {
	input := []int{1, 3, 15, 20, 4, 2}
	output := []int{1, 2, 3, 4, 15, 20}

	p := NewPriorityQueue(cmpMoreInt)
	for idx, prior := range input {
		p.Push(prior, idx)
	}

	// check
	for _, expPrior := range output {
		prior, _ := p.Pop()
		if expPrior != prior {
			t.Errorf("test for sort int failed %v != %v", expPrior, prior)
		}
	}
}
