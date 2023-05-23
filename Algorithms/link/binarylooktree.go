package link

func (n *TreeNode) Find(target int) *TreeNode {
	if n == nil {
		return nil
	}

	if target == n.Val {
		return n
	} else if target < n.Val {
		return n.Left.Find(target)
	} else {
		return n.Right.Find(target)
	}
}

func (n *TreeNode) Insert(v int) *TreeNode {
	if n == nil {
		return &TreeNode{
			Val: v,
		}
	}

	parent := n
	cur := n
	for cur != nil {
		if v < cur.Val {
			// 指针是值，这里是值传递
			parent = cur
			cur = cur.Left
		} else {
			parent = cur
			cur = cur.Right
		}
	}

	if v < parent.Val {
		parent.Left = &TreeNode{Val: v}
	} else {
		parent.Right = &TreeNode{Val: v}
	}

	return n
}
