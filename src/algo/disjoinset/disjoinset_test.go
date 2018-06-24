package disjoinset

import "testing"

func TestJoinSets(t *testing.T) {
	set1 := []int{1, 2, 3}
	set2 := []int{4, 5, 6}

	s := New(7)

	for i := 1; i <= 6; i++ {
		s.MakeSet(i)
	}

	for _, i := range set1 {
		s.Union(set1[0], i)
	}

	for _, i := range set2 {
		s.Union(set2[0], i)
	}

	// check that set1 is one set and set2 not in set1
	for _, i := range set1 {
		if s.Find(set1[0]) != s.Find(i) {
			t.Errorf("[%v] not in set %v", i, set1)
		}

		if s.Find(set2[0]) == s.Find(i) {
			t.Errorf("[%v] in set %v", i, set2)
		}
	}

	for _, i := range set2 {
		if s.Find(set2[0]) != s.Find(i) {
			t.Errorf("[%v] not in set %v", i, set2)
		}

		if s.Find(set1[0]) == s.Find(i) {
			t.Errorf("[%v] in set %v", i, set1)
		}
	}

	s.Union(5, 2)

	// check that all items in union set
	for _, i := range set1 {
		if s.Find(set1[0]) != s.Find(i) {
			t.Errorf("[%v] not in set %v union %v", i, set1, set2)
		}
	}

	for _, i := range set2 {
		if s.Find(set1[0]) != s.Find(i) {
			t.Errorf("[%v] not in set %v union %v", i, set1, set2)
		}
	}

}
