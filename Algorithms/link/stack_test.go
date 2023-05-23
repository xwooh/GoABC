package link

import "testing"

func TestPush(t *testing.T) {
	s := Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	if s.n != 3 {
		t.Errorf("Expected length of 3 but got %d", s.n)
	}
	t.Log("test stack push ok\n")
}

func TestPop(t *testing.T) {
	s := Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	val, _ := s.Pop()
	if val != 3 {
		t.Errorf("Expected pop value of 3 but got %v", val)
	}
	if s.n != 2 {
		t.Errorf("Expected length of 2 but got %d", s.n)
	}
	t.Log("test stack pop ok\n")
}

func TestTop(t *testing.T) {
	s := Stack{}
	s.Push(1)
	s.Push(2)
	s.Push(3)
	val, _ := s.Top()
	if val != 3 {
		t.Errorf("Expected top value of 3 but got %v", val)
	}
	if s.n != 3 {
		t.Errorf("Expected length of 3 but got %d", s.n)
	}
	t.Log("test stack top ok\n")
}
