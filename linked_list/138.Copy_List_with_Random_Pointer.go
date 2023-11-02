package main

type Node138 struct {
	Val    int
	Next   *Node138
	Random *Node138
}

func copyRandomList(head *Node138) *Node138 {
	if head == nil {
		return nil
	}

	copiedList := make(map[*Node138]*Node138)

	current := head
	for current != nil { // copiando os nós
		copiedList[current] = &Node138{Val: current.Val}
		current = current.Next
	}

	current = head
	for current != nil { // settando os pointers
		copiedList[current].Next = copiedList[current.Next]
		copiedList[current].Random = copiedList[current.Random]
		current = current.Next
	}

	return copiedList[head]

}
