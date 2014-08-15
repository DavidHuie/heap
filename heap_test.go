package heap

import (
	"reflect"
	"testing"
)

type hint int

func (h hint) Comp(i Interface) bool {
	return h < i.(hint)
}

func TestInsert(t *testing.T) {
	h := NewHeap()

	data := []hint{1, 3, 2, 4, 6, 2, 3}

	for _, i := range data {
		h.Insert(i)
	}

	values := make([]hint, 0)

	for {
		val, err := h.Delete()
		if err != nil {
			break
		} else {
			values = append(values, val.(hint))
		}
	}

	if !reflect.DeepEqual(values, []hint{6, 4, 3, 3, 2, 2, 1}) {
		t.Errorf("Invalid response: %v", values)
	}
}
