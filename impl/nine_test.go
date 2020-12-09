package impl

import "testing"

func TestFindSumPair(t *testing.T) {
	if x, y, ok := findSumPair([]int{1, 2, 3, 4, 5}, 7); !ok || x != 2 || y != 5 {
		t.Errorf("Failed! Expected (2,5), got (%v, %v)", x, y)
	}

	if _, _, ok := findSumPair([]int{1, 2, 3, 4, 5}, 200); ok {
		t.Errorf("Failed! Expected !ok")
	}
}
