package link

func invert(node *TreeNode) {
	// 翻转左右节点
	if node == nil {
		return
	}

	// 翻转当前节点的左右子节点
	tmp := node.Left
	node.Left = node.Right
	node.Right = tmp

	// 继续处理左右子节点
	invert(node.Left)
	invert(node.Right)
}

func InvertTree(node *TreeNode) {
	invert(node)
}
