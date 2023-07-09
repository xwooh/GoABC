package link

import "testing"


/*
        1
      /   \
     2     5
            \
             7
*/

var rg = &TreeNode{
	Val: 1,
	Left: &TreeNode{
		Val: 2,
	},
	Right: &TreeNode{
		Val: 5,
		Right: &TreeNode{
			Val:   7,
			Left:  nil,
			Right: nil,
		},
	},
}


func TestTreeMaxDepth(t *testing.T) {
	var d = MaxDepth(rg)
	if d != 3 {
		t.Errorf("calc tree's max depth failed => %d != 3", d)
	}
}

