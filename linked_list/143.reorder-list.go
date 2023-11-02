package main

// type ListNode struct {
// 	Val int
// 	Next *ListNode
// }

// PROBLEMA:
// You are given the head of a singly linked-list. The list can be represented as:

// L0 → L1 → … → Ln - 1 → Ln
// Reorder the list to be on the following form:

// L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
// Input: head = [1,2,3,4]
// Output: [1,4,2,3]

func reorderList(head *ListNode) {
	// precisamos achar o ponto do "meio" da lista
	slow, fast := head, head.Next
	// slow deve caminhar 1 e fast caminhar 2
	// precisamos dos dois ponteiros para garantir que não vai percorrer elementos a mais
	// se a lista tiver quantidade impar, o slow vai parar o loop, caso contrário será o fast

	for fast != nil && fast.Next != nil { // enquanto fast != nil e fast não chegou ao fim da lista
		slow = slow.Next      // slow vai ir pro pŕoximo valor
		fast = fast.Next.Next // fast vai ir pra 2 valores seguintes
	}
	// a partir daqui vamos ter o começo da segunda metade da lista

	// precisamos pegar o segundo ponteiro
	// o começo da segunda lista nesse ponto vai ser slow.Next
	second := slow.Next
	// vamos reverter a segunda lista
	// antes de reverter vamos quebrar esse link, pois já armazenamos no nosso ponteiro "second"
	slow.Next = nil
	var prev *ListNode // inicializada como nil
	// agora vamos fazer o revert da SEGUNDA PARTE DA LISTA

	for second != nil { // enquanto second não for nil, vamos setar para o próximo node
		tmp := second.Next // próximo elemento da lista
		// antes de fazer o shift do pointer vamos fazer o próximo elemento de second ser prev ( inicialmente nil )
		second.Next = prev
		prev = second
		second = tmp // second vai ser o próximo elemento da lista
	}

	// Fazendo o merge das duas metades das listas ( somente depois que a segunda metade da lista foi revertida)
	// precisamos achar o começo da segunda lista:
	first, second := head, prev // second vai ser nil depois do loop anterior, mas prev vai ser o último node da lista, que vai ser a nova head da segunda metade da lista
	for second != nil {         // a segunda lsita pode ser menor que a primeira, então enquanto o segundo ponteiro não for nil vai continuar o loop do merge das listas
		// precisamos de variaveis temporarias para quebrar os links
		tmp1, tmp2 := first.Next, second.Next
		first.Next = second
		second.Next = tmp1 // estamos inserindo o segundo node entre o first e o first.Next
		first = tmp1
		second = tmp2
	}

}
