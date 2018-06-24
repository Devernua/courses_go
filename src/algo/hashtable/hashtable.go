package hashtable

import (
	"container/list"
	"bytes"
)

// TODO: dynamically reallocate and calculate if need
const primeHashTable uint64 = 1000000007 // prime
const randomX uint32 = 263

type ValueType interface {
	Data() []byte
}

type HashTable struct {
	maxSize uint32
	arr     []list.List

	// cache for x^n
	cachePowersX []uint32
}

// Init initializes or clears hash table h.
func (h *HashTable) Init(size uint32) *HashTable {
	h.maxSize = size
	h.arr = make([]list.List, size)
	for _, l := range h.arr {
		l.Init()
	}

	return h
}

// New returns an initialized hash table.
func New(size uint32) *HashTable { return new(HashTable).Init(size) }

func (h *HashTable) HashFunction(value ValueType) (result uint32) {
	for idx, b := range value.Data() {
		// cache new power for x
		if idx >= len(h.cachePowersX) {
			if idx == 0 {
				h.cachePowersX = append(h.cachePowersX, 1)
			} else {
				newVal := uint64(randomX) * uint64(h.cachePowersX[idx-1])
				h.cachePowersX = append(h.cachePowersX, uint32(newVal%primeHashTable))
			}
		}

		newVal := uint64(result) + uint64(h.cachePowersX[idx])*(uint64)(b)
		result = uint32(newVal % primeHashTable)
	}

	return result % h.maxSize
}

func (h *HashTable) Insert(value ValueType) {
	if h.Find(value) == nil {
		foundList := &h.arr[h.HashFunction(value)]
		foundList.PushFront(value)
	}
}

func (h *HashTable) Delete(value ValueType) {
	foundList := &h.arr[h.HashFunction(value)]

	for v := foundList.Front(); v != nil; v = v.Next() {
		if bytes.Compare(value.Data(), v.Value.(ValueType).Data()) == 0 {
			foundList.Remove(v)
			return
		}
	}
}

func (h HashTable) Find(value ValueType) ValueType {
	idx := h.HashFunction(value)
	foundList := h.arr[idx]

	for v := foundList.Front(); v != nil; v = v.Next() {
		if bytes.Compare(value.Data(), v.Value.(ValueType).Data()) == 0 {
			return v.Value.(ValueType)
		}
	}

	return nil
}

func (h HashTable) Check(idx uint32) []ValueType {
	var result []ValueType
	for v := h.arr[idx].Front(); v != nil; v = v.Next() {
		result = append(result, v.Value.(ValueType))
	}

	return result
}
