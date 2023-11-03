package main

type ListNode141 struct {
	Val  int
	Next *ListNode141
}

func hasCycle(head *ListNode141) bool {
	// podemos criar um hash map / hash set para conferir se já passamos pelo elemento anteriormente
	visited := make(map[*ListNode141]*ListNode141)

	current := head
	for current != nil {
		if _, ok := visited[current]; !ok {
			visited[current] = &ListNode141{Val: current.Val}
		} else {
			return true
		}

		current = current.Next
	}

	return false
}
