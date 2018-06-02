package disjoinset

func NewDisJoinSet() DisJoinSet {
	s := DisJoinSet{}
	s.parent = map[int]int{}
	s.rank = map[int]int{}
	return s
}

type DisJoinSet struct {
	parent map[int]int
	rank map[int]int // there only root, dont forgot delete on union less, and update new root
}

func (s *DisJoinSet) MakeSet(idx int) {
	// TODO: check that not exist
	s.parent[idx] = idx
	s.rank[idx] = 0
}

func (s *DisJoinSet) Union(left, right int) {
	leftParent := s.Find(left)
	rightParent := s.Find(right)
	if leftParent == rightParent {
		// TODO: may be error?
		return
	}
	if s.rank[leftParent] > s.rank[rightParent] {
		s.parent[rightParent] = leftParent
		delete(s.rank, rightParent)
	} else {
		s.parent[leftParent] = rightParent
		delete(s.rank, leftParent)
		if s.rank[leftParent] == s.rank[rightParent] {
			s.rank[rightParent] += 1
		}
	}
}

func (s *DisJoinSet) Find(idx int) int {
	if idx != s.parent[idx] {
		s.parent[idx] = s.Find(s.parent[idx])
	}

	return s.parent[idx]
}