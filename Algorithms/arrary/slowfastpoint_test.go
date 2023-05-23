package arrary

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{1, 1, 2}

	c := RemoveDuplicates(nums)
	if c != 2 {
		t.Errorf("删除失败 %v != [1, 2, 2]", nums)
	}

	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	c = RemoveDuplicates(nums)
	if c != 5 {
		t.Errorf("删除失败 %v != [0, 1, 2, 3, 4, ...]", nums)
	}
}
