package link

import (
	"reflect"
	"testing"
)

/*
      1
    /   \
   2     5
  / \     \
 8   9     7

=>

      1
    /   \
   5     2
  /     / \
 7     9   8
*/

var it = &TreeNode{
	Val: 1,
	Left: &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 8},
		Right: &TreeNode{Val: 9},
	},
	Right: &TreeNode{
		Val:   5,
		Right: &TreeNode{Val: 7},
	},
}

func TestInvertTree(t *testing.T) {
	InvertTree(it)
	pt := PreTraverse(it)
	if !reflect.DeepEqual([]int{1, 5, 7, 2, 9, 8}, pt) {
		t.Errorf("反转二叉树错误: %v", pt)
	}
}
