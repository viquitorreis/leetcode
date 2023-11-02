package main

import "fmt"

type ListNode19 struct {
	Val  int
	Next *ListNode19
}

func buildList() *ListNode19 {
	node5 := &ListNode19{Val: 5, Next: nil}
	node4 := &ListNode19{Val: 4, Next: node5}
	node3 := &ListNode19{Val: 3, Next: node4}
	node2 := &ListNode19{Val: 2, Next: node3}
	node1 := &ListNode19{Val: 1, Next: node2}

	return node1
}

func removeNthFromEnd(head *ListNode19, n int) *ListNode19 {

	dummy := &ListNode19{Next: head}
	slow, fast := dummy, dummy
	fmt.Println("slow as dummy =>", slow)

	for i := 0; i <= n; i++ {
		fast = fast.Next
		fmt.Println("fast 1º loop =>", fast)
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	fmt.Println("slow 1º =>", slow)
	fmt.Println("fast 2º =>", fast)

	slow.Next = slow.Next.Next

	return dummy.Next

}
