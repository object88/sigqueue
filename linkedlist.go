package sigqueue

// LinkedList is an implimentation of a doubly-linked list with minimal
// functionality to support a Sigqueue.
type LinkedList struct {
	head *linkedListItem
	tail *linkedListItem
}

type linkedListItem struct {
	item IDer
	prev *linkedListItem
	next *linkedListItem
}

// CreateLinkedList returns a new LinkedList structure.
func CreateLinkedList() *LinkedList {
	return &LinkedList{nil, nil}
}

// Peek returns the item at the head of the linked list without popping it off,
// so the linked list is not changed.  If the linked list is empty or nil, it
// returns nil.
func (ll *LinkedList) Peek() IDer {
	if ll == nil || ll.head == nil {
		return nil
	}

	return ll.head.item
}

// Peer returns the item at the tail of the linked list without popping it off,
// so the linked list is not changed.  If the linked list is empty or nil, it
// returns nil.
func (ll *LinkedList) Peer() IDer {
	if ll == nil || ll.tail == nil {
		return nil
	}

	return ll.tail.item
}

// Push adds the item to the end of the list.  If the pointer receiver is nil,
// a new linked list is returned.  Otherwise, it returns the same linked list.
func (ll *LinkedList) Push(item IDer) *LinkedList {
	if ll == nil {
		ll = CreateLinkedList()
	}

	lli := &linkedListItem{item, nil, nil}

	if ll.tail == nil {
		ll.head = lli
		ll.tail = lli
	} else if ll.head == ll.tail {
		ll.tail = lli
		ll.head.next = lli
		ll.tail.prev = ll.head
	} else {
		lli.prev = ll.tail
		ll.tail.next = lli
		ll.tail = lli
	}

	return ll
}

// Pop returns the first item in the linked list, or nil.
func (ll *LinkedList) Pop() IDer {
	if ll == nil || ll.head == nil {
		return nil
	}

	lli := ll.head

	if ll.head == ll.tail {
		ll.head = nil
		ll.tail = nil
	} else {
		ll.head = ll.head.next
		if ll.head != nil {
			ll.head.prev = nil
		}
	}

	return lli.item
}
