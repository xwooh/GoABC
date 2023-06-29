package link

import "fmt"

func PreOrder(n *TreeNode, container []int) []int {
	// 根(R) -> 左(l) -> 右(r)
	if n == nil {
		return container
	}

	// 前序
	container = append(container, n.Val)
	fmt.Printf("visit %d\n", n.Val)

	// 中序
	container = PreOrder(n.Left, container)

	// 后序
	container = PreOrder(n.Right, container)

	return container
}

func InOrder(n *TreeNode, container []int) []int {
	// 左(l) -> 根(R) -> 右(r)
	if n != nil {
		container = InOrder(n.Left, container)
		container = append(container, n.Val)
		container = InOrder(n.Right, container)
	}
	return container
}

func PostOrder(n *TreeNode, container []int) []int {
	// 左(l) -> 右(r) -> 根(R)
	if n != nil {
		container = PostOrder(n.Left, container)
		container = PostOrder(n.Right, container)
		container = append(container, n.Val)
	}
	return container
}
