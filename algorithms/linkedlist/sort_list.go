func sortList(head *ListNode) *ListNode {
    return merge(head, nil)
}

func merge(head, tail *ListNode) *ListNode {
    if head == tail { return tail }
    slow, fast := head, head
    for fast != tail && fast.Next != tail {
        slow, fast = slow.Next, fast.Next.Next
    }
    next := slow.Next
    slow.Next = nil
    return mergeNode(merge(head, slow), merge(next, tail))
}

func mergeNode(head1, head2 *ListNode) *ListNode {
    dummy := &ListNode{}
    curr := dummy
    for head1 != nil && head2 != nil {
        if head1.Val < head2.Val {
            curr.Next = &ListNode{Val: head1.Val}
            curr = curr.Next
            head1 = head1.Next
        } else {
            curr.Next = &ListNode{Val: head2.Val}
            curr = curr.Next
            head2 = head2.Next
        }
    }
    for head1 != nil {
        curr.Next = &ListNode{Val: head1.Val}
        curr = curr.Next
        head1 = head1.Next
    }
    for head2 != nil {
        curr.Next = &ListNode{Val: head2.Val}
        curr = curr.Next
        head2 = head2.Next
    }
    return dummy.Next
}
