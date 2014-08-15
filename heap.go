package heap

import (
	"errors"
	"log"
)

type Interface interface {
	Less(i Interface) bool
}

type Heap struct {
	data []Interface
}

func NewHeap() *Heap {
	return &Heap{make([]Interface, 0)}
}

func parentIndex(i int) (int, error) {
	if i <= 0 {
		return 0, errors.New("invalid index")
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

func (h *Heap) Insert(e Interface) {
	h.data = append(h.data, e)
	index := len(h.data) - 1

	log.Printf("Inserting: %v", e)

	// Move element from base to correct positioning.
	for {
		parentI, err := parentIndex(index)
		if err != nil {
			break
		}

		log.Printf("Parent index: %v", parentI)

		// Swap with parent if necessary.
		if h.data[index].Less(h.data[parentI]) {
			break
		} else {
			log.Printf("Swapping: %v & %v", parentI, index)
			h.swap(index, parentI)
			index = parentI
		}
	}

	log.Printf("Result: %v", h.data)
}

func (h *Heap) Delete() (Interface, error) {
	dataLen := len(h.data)

	if dataLen == 0 {
		return nil, errors.New("empty heap")
	}

	value := h.data[0]

	// Truncate the data and make the last element the root.
	newRoot := h.data[dataLen-1]
	h.data[0] = newRoot
	h.data = h.data[:dataLen-1]
	dataLen = dataLen - 1

	log.Printf("Truncated: %v", h.data)

	index := 0
	var swap int

	// Trickle down root to correct position.
	for {
		c1, c2 := childIndicies(index)

		log.Printf("C1: %v", c1)
		log.Printf("C2: %v", c2)

		// Find correct child to attempt swap with.
		if c1 >= dataLen && c2 >= dataLen {
			break
		} else if c1 >= dataLen {
			swap = c2
		} else if c2 >= dataLen {
			swap = c1
		} else {
			if h.data[c1].Less(h.data[c2]) {
				swap = c2
			} else {
				swap = c1
			}
		}

		log.Printf("Current data: %v", h.data)
		log.Printf("Len: %v", dataLen)
		log.Printf("Index: %v", index)
		log.Printf("Swap: %v", swap)

		// Swap if necessary.
		if h.data[index].Less(h.data[swap]) {
			h.swap(index, swap)
			index = swap
		} else {
			break
		}
	}

	return value, nil
}
