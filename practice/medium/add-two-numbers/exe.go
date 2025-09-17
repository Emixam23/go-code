package add_two_numbers

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// Initialize the result list and a pointer to traverse it
	var lres = &ListNode{}
	var it = lres

	// Initialize carry to handle sums greater than 9
	var carry = 0

	// Loop until both lists are exhausted and there's no carry left
	for l1 != nil || l2 != nil || carry > 0 {
		// Calculate the sum of the current digits and the carry
		total := carry
		if l1 != nil {
			total += l1.Val
			l1 = l1.Next // Move to the next node in l1
		}
		if l2 != nil {
			total += l2.Val
			l2 = l2.Next // Move to the next node in l2
		}

		// Update carry for the next iteration
		carry = total / 10

		// Create a new node with the digit value (total % 10) and append it to the result list
		it.Next = &ListNode{Val: total % 10}
		it = it.Next
	}

	// Return the next node of the dummy head, which is the actual start of the result list
	// Note that lres just get garbage collected
	return lres.Next
}
