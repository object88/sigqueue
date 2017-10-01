package sigqueue

import (
	"fmt"
	"sync"
)

// Sigqueue is a signalling queue.
type Sigqueue struct {
	waiters     *LinkedList
	readied     *Heap
	waitersLock sync.Mutex
	readiedLock sync.Mutex
	notify      chan<- IDer
	signal      *Signal
}

// CreateSigqueue creates a new instance of Sigqueue
func CreateSigqueue(notify chan<- IDer) *Sigqueue {
	s := &Sigqueue{
		waiters: CreateLinkedList(),
		readied: CreateHeap(),
		notify:  notify,
		signal:  CreateSignal(),
	}

	go s.process()

	return s
}

// WaitOn insert an item to wait for, added at the bottom of the queue.  Items
// must be monotonically increasing by ID.  An error is returned if the
// ID regresses, and the item is not queued.
func (s *Sigqueue) WaitOn(item IDer) error {
	s.waitersLock.Lock()

	id := s.waiters.Peer()
	if id != nil && id.ID() >= item.ID() {
		s.waitersLock.Unlock()
		return NewErrOutOfOrderWait(item.ID())
	}

	s.waiters.Push(item)

	s.waitersLock.Unlock()

	return nil
}

// Ready signals that the item should be signalled when it is at the top
// of the queue.
func (s *Sigqueue) Ready(item IDer) error {
	id := item.ID()

	s.readiedLock.Lock()

	s.readied.Insert(id)
	s.signal.Signal()

	s.readiedLock.Unlock()

	return nil
}

func (s *Sigqueue) process() {
	for {
		select {
		case <-s.signal.Ready():
			s.waitersLock.Lock()
			s.readiedLock.Lock()

			for {
				id, err := s.readied.Minimum()
				if err != nil && err != ErrHeapEmpty {
					fmt.Printf("WE'RE GOING DOWN: %s\n", err.Error())
					break
				}
				item := s.waiters.Peek()
				if item == nil || item.ID() != id {
					// The signal received was not for the oldest wait
					break
				}

				// Remove the minimum value from the hash, and notify the consumer
				// that the oldest wait is ready.
				s.readied.RemoveMinimum()
				s.notify <- s.waiters.Pop()
			}

			s.readiedLock.Unlock()
			s.waitersLock.Unlock()
		}
	}
}
