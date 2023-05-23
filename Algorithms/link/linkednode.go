package link

/*
链表数据结构的实现

基本方法：
1. 实例化
2. 新增节点
3. 删除节点
4. 查询节点
5. 修改某个节点数据
*/

import (
	"fmt"
	"strings"
)

type Node struct {
	v    int
	next *Node
}

type LinkedNode struct {
	head *Node
}

func (node *LinkedNode) toString() string {
	if node.head == nil {
		return ""
	}

	cur := node.head
	var str strings.Builder
	str.WriteString(fmt.Sprintf("%d", cur.v))

	for cur.next != nil {
		cur = cur.next
		str.WriteString(fmt.Sprintf("->%d", cur.v))
	}

	return str.String()
}

// toReverseString 栈方式实现反向打印链表
func (node *LinkedNode) toReverseString() string {
	var stack []int

	cur := node.head
	for cur != nil {
		stack = append(stack, cur.v)
		cur = cur.next
	}

	var str strings.Builder
	for i := len(stack) - 1; i >= 0; i-- {
		if i == 0 {
			str.WriteString(fmt.Sprintf("%d", stack[i]))
		} else {
			str.WriteString(fmt.Sprintf("%d->", stack[i]))
		}
	}
	return str.String()
}

// reversePrintByRecursion 递归方式实现的反向打印链表
func reversePrintByRecursion(node *Node) {
	if node != nil {
		if node.next != nil {
			// 入栈
			reversePrintByRecursion(node.next)
		}
		// 出栈: 打印出栈时 node 的值
		fmt.Println(node.v)
	}
}

// Append 在尾结点后面增加一个新节点
// 比如 LinkedNode{nil} 新增一个值为 1 的节点之后变为 LinkedNode{&Node(1, nil)}
func (node *LinkedNode) Append(v int) *LinkedNode {
	fmt.Println("==================")
	newNode := &Node{v: v}

	head := node.head
	if head == nil {
		node.head = newNode
	} else {
		for printNodeIterInfo(head); head.next != nil; printNodeIterInfo(head) {
			head = head.next
		}
		head.next = newNode
	}

	printAddNewNodeInfo(v)
	return node
}

// DeleteByIndex TODO 删除指定某个索引的节点
func (node *LinkedNode) DeleteByIndex(index int) *LinkedNode {
	return node
}

// DeleteByValue 删除第一个含有某值的节点
func (node *LinkedNode) DeleteByValue(v int) *LinkedNode {
	/*
		遍历链表：
		1. 判断头结点
			- 头结点 nil，直接返回 node
			- 头结点有值 & 头结点值 = v，返回 node.head.next
		2. 判断 cur.next 是否有值 & cur.next.v ?= v
			- 不等，往后迭代
			- 相等，删除 cur.next，即：cur.next = cur.next.next
	*/

	if node.head == nil {
		return node
	}

	if node.head.v == v {
		node.head = node.head.next
		return node
	}

	cur := node.head
	for cur.next != nil && cur.next.v != v {
		// 只要还有下一个节点 & 下一个节点的值 != v，就不断往下遍历
		cur = cur.next
	}

	// 如果上面跳出循环的条件是找到了目标值，则删除找到的这个节点 (cur.next)
	if cur.next != nil && cur.next.v == v {
		cur.next = cur.next.next
	}

	return node
}

func printNodeIterInfo(node *Node) {
	fmt.Printf("iter to node: %v\n", node.v)
}

func printAddNewNodeInfo(v int) {
	fmt.Printf("add new node %v to LinkedNode\n", v)
}

func main() {
	node := &LinkedNode{
		head: &Node{
			v: 1,
			next: &Node{
				v: 2,
				next: &Node{
					v: 3,
					next: &Node{
						v: 4,
						next: &Node{
							v:    5,
							next: nil,
						},
					},
				},
			},
		},
	}

	fmt.Println(node.toString())
	fmt.Println(node.toReverseString())
	reversePrintByRecursion(node.head)
}
