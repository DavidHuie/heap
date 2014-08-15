package heap

import (
	"errors"
	"sync"
)

var (
	InvalidIndex = errors.New("invalid index")
	EmptyHeap    = errors.New("empty heap")
)

type Interface interface {
	Comp(i Interface) bool
}

type Heap struct {
	sync.Mutex
	data []Interface
}

func NewHeap() *Heap {
	return &Heap{
		data: make([]Interface, 0),
	}
}

func parentIndex(i int) (int, error) {
	if i <= 0 {
		return 0, InvalidIndex
	}
	return (i - 1) / 2, nil
}

func childIndicies(i int) (int, int) {
	return 2*i + 1, 2*i + 2
}

func (h *Heap) swap(i, j int) {
	a := h.data[i]
	h.data[i] = h.data[j]
	h.data[j] = a
}

// Inserts an element into the heap.
func (h *Heap) Insert(e Interface) {
	h.Lock()
	defer h.Unlock()

	h.data = append(h.data, e)
	index := len(h.data) - 1

	// Move element from base to correct positioning.
	for {
		parentI, err := parentIndex(index)
		if err != nil {
			break
		}

		// Swap with parent if necessary.
		if h.data[index].Comp(h.data[parentI]) {
			break
		} else {
			h.swap(index, parentI)
			index = parentI
		}
	}

}

// Deletes the root element from the heap and returns it.
func (h *Heap) Delete() (Interface, error) {
	h.Lock()
	defer h.Unlock()

	dataLen := len(h.data)

	if dataLen == 0 {
		return nil, EmptyHeap
	}

	root := h.data[0]

	// Truncate the data and make the last element the root.
	newRoot := h.data[dataLen-1]
	h.data[0] = newRoot
	h.data = h.data[:dataLen-1]
	dataLen = dataLen - 1

	index := 0
	var swap int

	// Trickle down root to correct position.
	for {
		c1, c2 := childIndicies(index)

		// Find correct child to attempt swap with.
		if c1 >= dataLen && c2 >= dataLen {
			break
		} else if c1 >= dataLen {
			swap = c2
		} else if c2 >= dataLen {
			swap = c1
		} else {
			if h.data[c1].Comp(h.data[c2]) {
				swap = c2
			} else {
				swap = c1
			}
		}

		// Swap if necessary.
		if h.data[index].Comp(h.data[swap]) {
			h.swap(index, swap)
			index = swap
		} else {
			break
		}
	}

	return root, nil
}
