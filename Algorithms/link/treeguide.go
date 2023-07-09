package link

func traverse(n *TreeNode, currD, maxD int) int {
	if n == nil {
		return maxD
	}

	currD++
	if n.Left == nil && n.Right == nil {
		if currD > maxD {
			maxD = currD
		}
	}

	maxD = traverse(n.Left, currD, maxD)
	maxD = traverse(n.Right, currD, maxD)

	currD--

	return maxD

}

func MaxDepth(t *TreeNode) int {
	// 返回 `TreeNode` 的最大深度

	var maxD int

	if t == nil {
		return maxD
	}

	maxD = traverse(t, 0, 0)

	return maxD
}
