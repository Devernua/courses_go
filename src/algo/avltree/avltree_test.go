package avltree

import "testing"

func TestAVLright(t *testing.T) {
	tree := New()
	input := []int{1, 2, 3, 4, 5}
	for _, i := range input {
		tree.Insert(i)
	}

	n := tree.head
	for _, i := range input {
		if i != n.value {
			t.Errorf("[%v] not eq %v", i, n.value)
		}
		n = n.right
	}
}
