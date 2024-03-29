package linkedlist

// node represents an element of the linked list.
type Node[E comparable] struct {
	Data E
	Next *Node[E]
}

// linkedlist represents collection of nodes
type Linkedlist[E comparable] struct {
	Head *Node[E]
	Tail *Node[E]
	Len  int
}

// NewLinkedList constructs an empty linked list
func NewList[E comparable]() *Linkedlist[E] {
	return new(Linkedlist[E])
}

// Add inserts node at the beginning of the list
func (l *Linkedlist[E]) Add(v E) {
	l.AddFirst(v)
}

// addFirst inserts node at the beginning of list
func (l *Linkedlist[E]) AddFirst(v E) {
	n := &Node[E]{Data: v}
	if l.Head == nil {
		l.Tail = n
	} else {
		n.Next = l.Head
	}
	l.Head = n
	l.Len++
}

// AddLast adds node at the end of list
func (l *Linkedlist[E]) AddLast(v E) {
	if l.Head == nil {
		l.AddFirst(v)
	} else {
		n := &Node[E]{Data: v}
		l.Tail.Next = n
		l.Tail = n
		l.Len++
	}
}

// AddBefore inserts node before specified node in list.
// before node must not be nil.
// List is not modified if before is nil or does not exist.
func (l *Linkedlist[E]) AddBefore(bf *Node[E], v E) {
	if bf == nil {
		return
	}

	if l.Head.Data == bf.Data {
		l.AddFirst(v)
		return
	}

	if l.Contains(bf) {
		p := l.findPrev(bf)
		if p != nil {
			n := &Node[E]{Data: v}
			n.Next = p.Next
			p.Next = n
			l.Len++
		}
	}
}

// AddAfter inserts node after specified node in list
func (l *Linkedlist[E]) AddAfter(af *Node[E], v E) {
	if af == nil || l.Head == nil {
		return
	}

	n := &Node[E]{Data: v}

	if l.Head.Data == af.Data {
		n.Next = l.Head.Next
		l.Head.Next = n
	} else if l.Tail.Data == af.Data {
		l.Tail.Next = n
		l.Tail = n
	} else {
		if l.Contains(af) {
			curr := l.Head
			for curr.Data != af.Data {
				curr = curr.Next
			}
			n.Next = curr.Next
			curr.Next = n
		}
	}

	l.Len++
}

// RemoveFirst deletes node from head of list
func (l *Linkedlist[E]) RemoveFirst() (out E) {
	if l.Head == nil {
		return out
	}
	out = l.Head.Data
	l.Head = l.Head.Next
	l.Len--
	return
}

// RemoveLast deletes node from end of list
func (l *Linkedlist[E]) RemoveLast() (out E) {
	if l.Head == nil {
		return out
	}

	var prev *Node[E]
	curr := l.Head
	for curr.Next != nil {
		prev = curr
		curr = curr.Next
	}

	var n *Node[E]
	if prev == nil {
		n = l.Head
		l.Head = nil
		l.Tail = nil
	} else {
		n = prev.Next
		prev.Next = prev.Next.Next
		l.Tail = prev
	}

	l.Len--
	out = n.Data
	return
}

// RemoveAt deletes a node at specified index
func (l *Linkedlist[E]) RemoveAt(index int) (out E) {
	if l.Head == nil || index < 0 || index > l.Len {
		return out
	}

	var prev *Node[E]
	curr := l.Head

	for i := 0; i < index-1; i++ {
		prev = curr
		curr = curr.Next
	}

	if prev == nil {
		l.Head = curr.Next
	} else {
		prev.Next = curr.Next
	}

	if curr == l.Tail {
		l.Tail = prev
	}

	out = curr.Data
	l.Len--

	return
}

// GetAt retrieves and returns node at specified index
func (l *Linkedlist[E]) GetAt(index int) (out E) {
	if index < 0 || index > l.Len {
		return out
	}

	curr := l.Head
	for i := 0; i < index-1; i++ {
		curr = curr.Next
	}

	out = curr.Data
	return
}

// Contains checks if node exists in list
func (l *Linkedlist[E]) Contains(n *Node[E]) bool {
	curr := l.Head
	for curr != nil {
		if curr.Data == n.Data {
			return true
		}
		curr = curr.Next
	}
	return false
}

// findPrev returns node which preceeds specified node
func (l *Linkedlist[E]) findPrev(n *Node[E]) *Node[E] {
	var prev *Node[E]
	curr := l.Head
	for curr != nil && curr.Data != n.Data {
		prev = curr
		curr = curr.Next
	}
	return prev
}
