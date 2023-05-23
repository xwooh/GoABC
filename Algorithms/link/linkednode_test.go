package link

import (
	"testing"
)

func TestAppend(t *testing.T) {
	ln := LinkedNode{}

	rangeValues := make([]struct{}, 5)
	for idx := range rangeValues {
		ln.Append(idx + 1)
	}

	// 数据是否新增
	var increased int
	head := ln.head
	for head != nil {
		increased++
		head = head.next
	}
	if increased != len(rangeValues) {
		t.Errorf("linked node append new value failed")
	}
	t.Log("test linked node append ok\n")
}

func TestDeleteByValue(t *testing.T) {
	// create a linked list with values 1, 2, 3, 4, 5
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

	// delete a value that exists in the linked list
	node.DeleteByValue(3)
	expected := "1->2->4->5"
	if result := node.toString(); result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	// delete a value that doesn't exist in the linked list
	node.DeleteByValue(6)
	expected = "1->2->4->5"
	if result := node.toString(); result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	// delete the head node
	node.DeleteByValue(1)
	expected = "2->4->5"
	if result := node.toString(); result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	// delete the tail node
	node.DeleteByValue(5)
	expected = "2->4"
	if result := node.toString(); result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	// delete the only node in the linked list
	node.DeleteByValue(2)
	node.DeleteByValue(4)
	expected = ""
	if result := node.toString(); result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}

	// delete from empty node
	node.DeleteByValue(1)
	expected = ""
	if result := node.toString(); result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
	t.Log("test linked node DeleteByValue ok\n")
}
