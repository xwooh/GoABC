/*
简单起见，将算术表达式简化为只包含加减乘除四则运算，比如：`34 + 13 * 9 + 44 - 12 / 5`。

主要思路：
1. 我们从左向右遍历表达式，当遇到数字，压入数字栈
2. 遇到运算符就与运算符栈的栈顶元素进行比较
	1. 如果当前运算符比栈顶运算符优先级高，压入栈
	2. 如果当前运算符比栈顶运算符优先级低或相同，就从数字栈取出数字，从运算符栈取出运算符进行计算，计算完结果后，数字压入数字栈，当前运算符再压入运算符栈
*/

package stack

import (
	"fmt"
	"strings"
)

func calc(n0, n1 float32, o string) float32 {
	var ret float32
	switch o {
	case "+":
		ret = n0 + n1
	case "-":
		ret = n0 - n1
	case "*":
		ret = n0 * n1
	case "/":
		ret = n0 / n1
	default:
		fmt.Printf("计算失败, 运算符不合法: %v", o)
		return 0
	}
	return ret
}

type Calculator struct {
	ns []float32
	os []string
}

func (c *Calculator) AddNumber(n float32) {
	c.ns = append(c.ns, n)
}

func (c *Calculator) AddOperator(o string) {
	if len(c.os) == 0 || len(c.ns) < 1 {
		// 不满足计算条件，直接加进去
		c.os = append(c.os, o)
	} else {
		// 和当前栈顶的运算符进行比较
		topO := c.os[len(c.os)-1]
		if strings.ContainsAny(topO, "*/") || strings.ContainsAny(o, "+-") {
			// 当前运算符的优先级肯定小于等于此时栈顶的运算符，那就停一下，取出这个栈顶运算符进行计算
			c.os = c.os[:len(c.os)-1]

			// 取出两个数字
			n1 := c.ns[len(c.ns)-1]
			n0 := c.ns[len(c.ns)-2]
			// 弹出两个数字
			c.ns = c.ns[:len(c.ns)-2]

			ret := calc(n0, n1, topO)

			// 将算出来的数压入数字栈
			c.ns = append(c.ns, ret)
			// 将运算符压入运算符栈
			c.os = append(c.os, o)
		} else {
			// 当前运算符要比栈顶的优先级高，直接入栈
			c.os = append(c.os, o)
		}
	}
}

func (c *Calculator) Calc() float32 {
	// 此时运算符栈内的元素由上到下肯定是优先级由高到低，所以依次取出来计算即可
	for len(c.os) > 0 {
		if len(c.ns) < 2 {
			fmt.Printf("计算失败, 参与计算的数字不足")
			return 0
		}

		topO := c.os[len(c.os)-1]
		// 弹出栈顶运算符
		c.os = c.os[:len(c.os)-1]

		// 取出两个数字
		n1 := c.ns[len(c.ns)-1]
		n0 := c.ns[len(c.ns)-2]
		// 弹出两个数字
		c.ns = c.ns[:len(c.ns)-2]

		ret := calc(n0, n1, topO)
		c.ns = append(c.ns, ret)
	}
	if len(c.ns) == 1 {
		return c.ns[0]
	}
	fmt.Printf("计算失败, 参与的运算符不足")
	return 0
}

func main() {
	c := Calculator{}

	c.AddNumber(34)
	c.AddOperator("+")
	c.AddNumber(13)
	c.AddOperator("*")
	c.AddNumber(9)
	c.AddOperator("+")
	c.AddNumber(44)
	c.AddOperator("-")
	c.AddNumber(12)
	c.AddOperator("/")
	c.AddNumber(5)

	ret := c.Calc()
	if ret != 192.6 {
		fmt.Printf("calculate error, %v != 192.6\n", ret)
	} else {
		fmt.Printf("计算成功！\n")
	}
}
