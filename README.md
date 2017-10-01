# sigqueue

A `Sigqueue` is a signalling queue which maintains two data structures, and signals when the first of each are matching.

## Supporting data structures

The `Sigqueue` depends on two data structures internally to track state.  On the incoming side, there is a monotonically increasing linked list of `IDer`s.  On the waiting side, there is a heap of ready ids.

### heap

A `Heap` is a standard [min-heap data structure](https://en.wikipedia.org/wiki/Heap_(data_structure)), typed specifically to support the `Sigqueue`.

### linkedlist

A `LinkedList` is a standard double-linked list, purpose built for `sigqueue`.

### signal

A `Signal` is a way to notify a consumer when a signal has been raised over channels.
