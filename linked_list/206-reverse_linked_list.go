package main

import "fmt"

// PROBLEMA
// Given the head of a singly linked list, reverse the list, and return the reversed list.
// Input: head = [1,2,3,4,5]
// Output: [5,4,3,2,1]

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var revHead *ListNode

	for head != nil {
		tmp := head.Next
		head.Next = revHead // aponta para o anterior
		// revHead == nil
		revHead = head
		fmt.Printf("revHead.Val => %d\n", revHead.Val)
		fmt.Printf("head => %d\n", head.Val)
		head = tmp
	}
	return revHead
}

func sliceToLinkedList(vals []int) *ListNode {
	var head, tail *ListNode

	for _, val := range vals {
		node := &ListNode{Val: val}
		if head == nil {
			head = node
			tail = node
		} else {
			tail.Next = node
			tail = node
		}
	}

	return head
}

func printLinkedList(head *ListNode) {
	current := head

	for current != nil {
		fmt.Printf("%d => ", current.Val)
		current = current.Next
	}

	fmt.Println("nil")
}
