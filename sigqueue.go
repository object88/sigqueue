package sigqueue

import (
	"fmt"
	"sync"
)

// Sigqueue is a signalling queue.
type Sigqueue struct {
	ll *LinkedList
	h  *Heap
	m  sync.Mutex
}

// CreateSigqueue creates a new instance of Sigqueue
func CreateSigqueue() *Sigqueue {
	return &Sigqueue{
		ll: CreateLinkedList(),
		h:  CreateHeap(),
	}
}

// WaitOn insert an item to wait for,  at the bottom of the queue.  Items
// must be monotonically increasing by ID.  An error is returned if the
// ID regresses, and the item is not queued.
func (s *Sigqueue) WaitOn(item IDer) error {
	s.m.Lock()

	id := s.ll.Peer()
	if id != nil && id.ID() >= item.ID() {
		s.m.Unlock()
		return fmt.Errorf("Cannot queue item; ID is not increasing")
	}

	s.ll.Push(item)

	s.m.Unlock()
	return nil
}

// Ready signals that the item should be signalled when it is at the top
// of the queue.
func (s *Sigqueue) Ready(item IDer) error {
	return nil
}
