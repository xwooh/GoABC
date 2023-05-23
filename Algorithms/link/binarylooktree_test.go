package link

import (
	"fmt"
	"reflect"
	"testing"
)

/*
	      33
	  /       \
	 17        50
	/ \     /    \
   13   18  34   58
	\    \     /  \
	 16  25   51  66
	    /  \
	   19  27
*/

var r2 = &TreeNode{
	Val: 33,
	Left: &TreeNode{
		Val: 17,
		Left: &TreeNode{
			Val: 13,
			Right: &TreeNode{
				Val: 16,
			},
		},
		Right: &TreeNode{
			Val: 18,
			Right: &TreeNode{
				Val: 25,
				Left: &TreeNode{
					Val: 19,
				},
				Right: &TreeNode{
					Val: 27,
				},
			},
		},
	},
	Right: &TreeNode{
		Val: 50,
		Left: &TreeNode{
			Val: 34,
		},
		Right: &TreeNode{
			Val: 58,
			Left: &TreeNode{
				Val: 51,
			},
			Right: &TreeNode{
				Val: 66,
			},
		},
	},
}

func TestBinaryLookTreeFind(t *testing.T) {
	n := r2.Find(13)
	if n.Val != 13 && n.Right.Val != 16 {
		t.Errorf("二叉查找树查找失败: 13 != %d", n.Val)
	}

	n2 := r2.Find(66)
	if n2.Val != 66 && n.Right != nil && n.Left != nil {
		t.Errorf("二叉查找树查找失败: 66 != %d", n.Val)
	}

	n3 := r2.Find(100)
	if n3 != nil {
		t.Errorf("二叉查找树查找失败，不应该查到值！")
	}
}

func TestBinaryLookTreeInsert(t *testing.T) {
	x := r2.Insert(55)

	var c []int
	c = InOrder(x, c)
	if !reflect.DeepEqual([]int{13, 16, 17, 18, 19, 25, 27, 33, 34, 50, 51, 55, 58, 66}, c) {
		t.Errorf("二叉查找树插入数据错误: %v", c)
	}
}

func TestX(t *testing.T) {
	m := map[int]string{
		1: "a",
		2: "b",
	}

	for a, b := range m {
		fmt.Printf("a=%v, b=%v\n", a, b)
	}
}
