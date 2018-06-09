package hashtable

import (
	"testing"
)

type stringWrapper struct {
	Str string
}

func (str *stringWrapper) Data() (buf []byte) {
	// TODO: optimize
	for _, c := range str.Str {
		buf = append(buf, byte(c))
	}
	return buf
}

func TestHashTable(t *testing.T) {
	h := New(5)
	data := []string{"world", "HellO"}
	for _, d := range data {
		h.Insert(&stringWrapper{d})
	}

	res := h.Check(4)
	if len(res) != 2 {
		t.Errorf("test check hash failed")
	}

	for idx, s := range res {
		curStr := s.(*stringWrapper).Str
		if curStr != data[len(res) - idx - 1] {
			t.Errorf("test check hash table failed %v != %v", curStr, data[idx])
		}
	}

	findResult := h.Find(&stringWrapper{"HellO"})
	if findResult == nil {
		t.Errorf("test check hash table find return nil")
	}

}