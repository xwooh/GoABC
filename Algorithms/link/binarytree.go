package link

import "fmt"

func PreOrder(n *TreeNode, container []int) []int {
	// 根(R) -> 左(l) -> 右(r)
	if n == nil {
		return container
	}

	// 根
	container = append(container, n.Val)
	fmt.Printf("visit %d\n", n.Val)

	// 左
	container = PreOrder(n.Left, container)

	// 右
	container = PreOrder(n.Right, container)

	return container
}

func PreTraverse(n *TreeNode) []int {
	var res []int

	if n == nil {
		return res
	}

	// 根
	res = append(res, n.Val)

	// 左
	res = append(res, PreTraverse(n.Left)...)

	// 右
	res = append(res, PreTraverse(n.Right)...)

	return res
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
