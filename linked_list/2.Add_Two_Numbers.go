package main

type ListNode2 struct {
	Val int
    Next *ListNode2
}
	*/
func addTwoNumbers(l1 *ListNode2, l2 *ListNode2) *ListNode2 {
	result := &ListNode2{}
	tmp := result // variavel temporária

	for l1 != nil || l2 != nil {
		if l1 != nil {
			tmp.Val += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			tmp.Val += l2.Val
			l2 = l2.Next
		}

		// transporte ou carry
		if tmp.Val > 9 {
			tmp.Val -= 10
			tmp.Next = &ListNode2{Val: 1} // depois de subtrair 10, temos que adicionar 1 no ṕŕoximo nó como valor de transporte
			// esse valor 1 vai ser passado para a PPRÓXIMA SOMA quando passaromos para os próximos digitos 
		} else if l1 != nil || l2 != nil {
			// caso tmp.Val < 9 significa que não precisa de transporte, mas se ainda tiver algum digito em l1 ou l2
			// vamos criar um próximo nó vazio para continuar a soma
			tmp.Next = &ListNode2{}
		}

		tmp = tmp.Next

	}

	return result
}