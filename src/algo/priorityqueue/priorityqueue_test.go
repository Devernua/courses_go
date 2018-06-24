package priorityqueue

import (
	"testing"
)

func cmpLessInt(l, r ValueType) bool {
	return l.(int) < r.(int)
}

func cmpMoreInt(l, r ValueType) bool {
	return r.(int) < l.(int)
}

func TestSortInt(t *testing.T) {
	input := []int{1, 3, 15, 20, 4, 2}
	output := []int{1, 2, 3, 4, 15, 20}

	p := New(cmpMoreInt)
	for _, prior := range input {
		p.Push(prior)
	}

	// check
	for _, expPrior := range output {
		prior := p.Pop()
		if expPrior != prior {
			t.Errorf("test for sort int failed %v != %v", expPrior, prior)
		}
	}
}
