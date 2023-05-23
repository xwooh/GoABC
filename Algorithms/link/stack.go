/*
栈数据结构的实现

基本方法：
1. 实例化
2. 入栈
3. 出栈
*/

package link

type Stack struct {
	n int
	c []int
}

func (s *Stack) Push(v int) {
	s.c = append(s.c, v)
	s.n += 1
}

func (s *Stack) Pop() (int, bool) {
	if s.n <= 0 {
		return 0, false
	}
	pv := s.c[s.n-1]
	s.c = s.c[:s.n-1]
	s.n -= 1
	return pv, true
}

func (s *Stack) Top() (int, bool) {
	if s.n <= 0 {
		return 0, false
	}
	return s.c[s.n-1], true
}
