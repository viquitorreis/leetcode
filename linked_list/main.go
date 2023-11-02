package main

import "fmt"

func main() {
	// headValues := []int{5, 4, 3, 2, 1}
	// head := sliceToLinkedList(headValues)

	// rvList := reverseList(head)
	// printLinkedList(rvList)

	// 19
	head := buildList()
	removedHead := removeNthFromEnd(head, 2)

	current := removedHead
	for current != nil {
		fmt.Printf("%d -> ", current.Val)
		current = current.Next
	}
	fmt.Println("nil")
}
