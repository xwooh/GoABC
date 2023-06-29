package link

import (
	"reflect"
	"testing"
)

/*
        1
      /   \
     2     5
    / \   / \
   3   4 6   7
*/

var r1 = &TreeNode{
	Val: 1,
	Left: &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   4,
			Left:  nil,
			Right: nil,
		},
	},
	Right: &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   6,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   7,
			Left:  nil,
			Right: nil,
		},
	},
}

func TestTreePreOrder(t *testing.T) {
	var c []int
	c = PreOrder(r1, c)
	if !reflect.DeepEqual([]int{1, 2, 3, 4, 5, 6, 7}, c) {
		t.Errorf("前序遍历错误: %v", c)
		t.Log("================>>>>>>>>>>>>>>")
	}
}

func TestTreeInOrder(t *testing.T) {
	var c []int
	c = InOrder(r1, c)
	if !reflect.DeepEqual([]int{3, 2, 4, 1, 6, 5, 7}, c) {
		t.Errorf("中序遍历错误: %v", c)
	}
}

func TestTreePostOrder(t *testing.T) {
	var c []int
	c = PostOrder(r1, c)
	if !reflect.DeepEqual([]int{3, 4, 2, 6, 7, 5, 1}, c) {
		t.Errorf("后序遍历错误: %v", c)
	}
}
